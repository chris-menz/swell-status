package middleware

import (
	"net/http"
	_ "github.com/lib/pq"
	"api/models"
	"encoding/json"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)



func CreateUser(w http.ResponseWriter, r *http.Request)  {
	var newuser models.User
	var userid int64

	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		log.Fatal(err)
	}

	missingParameters := 
		newuser.Username == "" ||
		newuser.FirstName == "" ||
		newuser.LastName == "" ||
		newuser.Email == ""

	if missingParameters {
		res := Response {
			Message: "Missing parameters",
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if len(newuser.HashedPassword) < 5 {
		failure_res := Response{
			Message: "Invalid Password",
		}
		err = json.NewEncoder(w).Encode(failure_res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	
	db := PostgresConnection()
	defer db.Close()

	//check if a user with this email or username is already registered
	selectUserQuery := "SELECT username FROM users WHERE username = $1"

	err = db.QueryRow(selectUserQuery, newuser.Username).Scan(&newuser.Username)

	if err != nil && err != sql.ErrNoRows {
        log.Fatal(err)
    } 
	// err being nil indicates that a row was found
	if err == nil {
		failure_res := Response{
			Message: "Username already registered",
		}
		err = json.NewEncoder(w).Encode(failure_res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	selectUserQuery = "SELECT email FROM users WHERE email = $1"

	err = db.QueryRow(selectUserQuery, newuser.Email).Scan(&newuser.Email)

	if err != nil && err != sql.ErrNoRows{
        log.Fatal(err)
    } 
	// err being nil indicates that a row was found
	if err == nil {
		failure_res := Response{
			Message: "Email already registered",
		}
		err = json.NewEncoder(w).Encode(failure_res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// if no user was found, then err == sql.ErrNoRows at this point in the code
	err = nil

	bytes, err := bcrypt.GenerateFromPassword([]byte(newuser.HashedPassword), 10)
	hashedPW := string(bytes)
	
	insertUserQuery := `INSERT INTO users 
	(username, first_name, last_name, email, hashed_password) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id`

	err = db.QueryRow(insertUserQuery, newuser.Username, newuser.FirstName, newuser.LastName, newuser.Email, hashedPW).Scan(&userid)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	success_res := Response{
		ID: int64(userid),
		Message: "User registered successfully",
	}

	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var user models.UserResponse

	db := PostgresConnection()
	defer db.Close()

	selectUserQuery := "SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u WHERE id=$1"
	row := db.QueryRow(selectUserQuery, id)

	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		failure_res := Response {
			Message: "No user found",
		}
		err = json.NewEncoder(w).Encode(failure_res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db := PostgresConnection()
	defer db.Close()

	getUsersQuery := `SELECT u.id, u.first_name, u.last_name, u.email, u.username FROM users u`

	rows, err := db.Query(getUsersQuery)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username)
		if err == sql.ErrNoRows {
			failure_res := Response {
				Message: "No surfboards found",
			}

			err = json.NewEncoder(w).Encode(failure_res)
			if err != nil {
				log.Fatal(err)
			}

			return
		}
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

