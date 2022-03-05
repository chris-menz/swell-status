package middleware

import(
	"github.com/gorilla/sessions"
	"net/http"
	_ "github.com/lib/pq"
	"api/models"
	"encoding/json"
	"log"
	"database/sql"
	"os"
	"golang.org/x/crypto/bcrypt"
)

// CHANGE TO ENV
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AuthResponse struct {
	Message string `json:"message,omitempty"`
	User models.UserResponse `json:"user,omitempty"`
}

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		if session.Values["userID"] != nil {
			next.ServeHTTP(w, r)
		} else{
			res := Response {
				Message: "Not authenticated",
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	})
}

func GetCurrUserID(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	res := Response {
		ID: session.Values["userID"].(int64),
		Message: "user is logged in",
	}
	json.NewEncoder(w).Encode(res)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var user models.UserResponse
	var passwordHash string

	session, _ := store.Get(r, "session")

	if session.Values["userID"] != nil {
		res := Response{
			Message: "already logged in",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Fatal(err)
	}
	
	db := PostgresConnection()
	defer db.Close()


	authUserQuery := `SELECT u.id, u.username, u.first_name, u.last_name, u.email, u.hashed_password FROM users u WHERE username=$1`
	err = db.QueryRow(authUserQuery, creds.Username).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &passwordHash)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		failure_res := Response{
			Message: "No user found with those credentials",
		}
		json.NewEncoder(w).Encode(failure_res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(creds.Password))
	if err != nil {
		failure_res := Response{
			Message: "No user found with those credentials",
		}
		json.NewEncoder(w).Encode(failure_res)
		return
	}

	session.Values["userID"] = user.ID
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure: true, 
		SameSite: http.SameSiteNoneMode,
	}
	
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res AuthResponse

	res = AuthResponse {
		Message: "login successful",
		User: user,
	}
	
	json.NewEncoder(w).Encode(res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure: true, 
		SameSite: http.SameSiteNoneMode,
	}
	delete(session.Values, "userID")

	err := session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}


	type SessionResponse struct {
		Message string `json:"message,omitempty"`
		SessionId int64 `json:"session_id,omitempty"`
	}

	res := SessionResponse {
		Message: "logged out",
	}

	json.NewEncoder(w).Encode(res)
}
