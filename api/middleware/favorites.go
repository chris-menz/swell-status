package middleware

import (
	"net/http"
	"api/models"
	_ "github.com/lib/pq"
	"encoding/json"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
)

func GetAllFavorites(w http.ResponseWriter, r *http.Request){
	db := PostgresConnection()
	defer db.Close()

	var favorites []models.Favorite = []models.Favorite{}

	getFavoritesQuery := "SELECT f.* from FAVORITES f"

	rows, err := db.Query(getFavoritesQuery)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var favorite models.Favorite

		err = rows.Scan(&favorite.ID, &favorite.SurfSessionID, &favorite.UserID)
		if err == sql.ErrNoRows {
			failure_res := Response {
				Message: "No favorites found",
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
		favorites = append(favorites, favorite)
	}

	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		log.Fatal(err)
	}
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	var newfavorite models.Favorite
	var favoriteid int64

	session, _ := store.Get(r, "session")

	err := json.NewDecoder(r.Body).Decode(&newfavorite)
	if err != nil {
		log.Fatal(err)
	}

	db := PostgresConnection()
	defer db.Close()

	insertFavoriteQuery := `INSERT INTO favorites (surf_session_id, user_id)
		VALUES ($1, $2)
		RETURNING id`

	err = db.QueryRow(insertFavoriteQuery, newfavorite.SurfSessionID, session.Values["userID"].(int64)).Scan(&favoriteid)
	if err != nil {
		log.Fatal(err)
	}

	success_res := Response{
		Message: "Favorite added",
	}
	
	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var userID int

	session, _ := store.Get(r, "session")
	
	db := PostgresConnection()
	defer db.Close()

	err := db.QueryRow(`SELECT f.user_id FROM favorites f WHERE id = $1`, id).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find favorite",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	if int64(userID) != session.Values["userID"].(int64) {
		res := Response{
			Message: "can't delete anothers favorite",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	deleteFavoriteQuery := `DELETE FROM favorites WHERE id=$1`
	sqlresp, err := db.Exec(deleteFavoriteQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "favorite " + id + " deleted",	
	}

	json.NewEncoder(w).Encode(res)
}