package router

import (
	"github.com/gorilla/mux"
	// "github.com/gorilla/csrf"
	"api/middleware"
	"net/http"
)

func Handlers() *mux.Router{
	r := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	// csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.TrustedOrigins([]string{"www.swellstatus.com"}))
	// r.Use(csrfMiddleware)
	r.Use(Middleware)

	//conditions
	r.HandleFunc("/weather/{lat}/{lng}", middleware.WeatherHandler).Methods("GET")
	r.HandleFunc("/owm/{lat}/{lng}", middleware.OWMHandler).Methods("GET")
	r.HandleFunc("/sg/{lat}/{lng}", middleware.SgHandler).Methods("GET")
	r.HandleFunc("/sg/{lat}/{lng}/{start}/{end}", middleware.PastSgHandler).Methods("GET")
	r.HandleFunc("/timezone/{lat}/{lng}", middleware.GetTimeZoneOffset).Methods("GET")
	r.HandleFunc("/tide/{lat}/{lng}/{date}", middleware.TideHandler).Methods("GET")

	//register + login
	r.HandleFunc("/createuser", middleware.CreateUser).Methods("POST")
	r.HandleFunc("/login", middleware.Login).Methods("POST")

	//get users
	r.HandleFunc("/user", middleware.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.GetUserByID).Methods("GET")

	//get surfboards
	r.HandleFunc("/surfboard/user/{user_id}", middleware.GetSurfboardsByUserID).Methods("GET")
	r.HandleFunc("/surfboard/{id}", middleware.GetSurfboardByID).Methods("GET")

	//get surf sessions
	r.HandleFunc("/surfsession", middleware.GetPublicSurfSessions).Methods("GET")

	//get comments
	r.HandleFunc("/comment", middleware.GetAllComments).Methods("GET")

	//get favorites
	r.HandleFunc("/favorite", middleware.GetAllFavorites).Methods("GET")	
	
	//get saved spots
	r.HandleFunc("/spot/user/{user_id}", middleware.GetSpotsByUserID).Methods("GET")
	
	//protected routes
	s := r.PathPrefix("/auth").Subrouter().StrictSlash(true)
	// s.Use(Middleware)
	s.Use(middleware.CheckAuth)

	s.HandleFunc("/getcurruserid", middleware.GetCurrUserID).Methods("GET")
	s.HandleFunc("/logout", middleware.Logout).Methods("POST")

	//surfboard crud
	s.HandleFunc("/surfboard/{id}", middleware.EditSurfboard).Methods("PUT")
	s.HandleFunc("/surfboard/{id}", middleware.DeleteSurfboard).Methods("DELETE")
	s.HandleFunc("/surfboard", middleware.CreateSurfboard).Methods("POST")

	//spots crud
	s.HandleFunc("/spot", middleware.AddSpot).Methods("POST")
	s.HandleFunc("/spot/{id}", middleware.DeleteSpot).Methods("DELETE")


	// surf sessions crud
	s.HandleFunc("/surfsession", middleware.CreateSurfSession).Methods("POST")
	s.HandleFunc("/surfsession/user/{user_id}", middleware.GetSurfSessionsByUserId).Methods("GET")
	s.HandleFunc("/surfsession/{id}", middleware.RemoveSurfSession).Methods("DELETE")

	//comments crud
	s.HandleFunc("/comment", middleware.AddComment).Methods("POST")
	s.HandleFunc("/comment/{id}", middleware.DeleteComment).Methods("DELETE")

	//favorites crud
	s.HandleFunc("/favorite", middleware.AddFavorite).Methods("POST")
	s.HandleFunc("/favorite/{id}", middleware.DeleteFavorite).Methods("DELETE")

	return r
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Origin", "https://condescending-hoover-f550c9.netlify.app")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
