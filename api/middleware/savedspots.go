package middleware

import(
	"net/http"
	"log"
	"encoding/json"
	"api/models"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/gorilla/mux"
)

func AddSpot(w http.ResponseWriter, r *http.Request) {
	var newSpot models.SavedSurfSpot
	var savedSpotID int64

	session, _ := store.Get(r, "session")

	err := json.NewDecoder(r.Body).Decode(&newSpot)
	if err != nil {
		log.Fatal(err)
	}

	if newSpot.SpotName == "" || newSpot.SpotRegion == "" {
		res := Response {
			Message: "Missing parameters",
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	db := PostgresConnection()
	defer db.Close()

	insertSpotQuery := `INSERT INTO saved_surf_spots (spot_name, spot_region, user_id)
		VALUES ($1, $2, $3)
		RETURNING id`

	err = db.QueryRow(insertSpotQuery, newSpot.SpotName, newSpot.SpotRegion, session.Values["userID"].(int64)).Scan(&savedSpotID)
	if err != nil {
		log.Fatal(err)
	}

	success_res := Response{
		Message: "spot saved",
	}
	
	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteSpot (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var spotUserID int

	session, _ := store.Get(r, "session")
	
	db := PostgresConnection()
	defer db.Close()

	err := db.QueryRow(`SELECT s.user_id FROM saved_surf_spots s WHERE id = $1`, id).Scan(&spotUserID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find spot",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	if int64(spotUserID) != session.Values["userID"].(int64) {
		res := Response{
			Message: "can't remove anothers spot",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	deleteSpotQuery := `DELETE FROM saved_surf_spots WHERE id=$1`
	sqlresp, err := db.Exec(deleteSpotQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "spot removed",	
	}

	json.NewEncoder(w).Encode(res)
}

func GetSpotsByUserID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	user_id := params["user_id"]
	var spots []models.SavedSurfSpot

	if user_id == "undefined" {
		failure_res := Response {
			Message: "Not logged in",
		}
		json.NewEncoder(w).Encode(failure_res)
		return
	}

	db := PostgresConnection()
	defer db.Close()

	getSpotsQuery := `SELECT s.id, s.spot_name, s.spot_region, s.user_id
		FROM saved_surf_spots s INNER JOIN users u ON s.user_id = u.id 
		WHERE s.user_id = $1`

	rows, err := db.Query(getSpotsQuery, user_id)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var spot models.SavedSurfSpot

		err = rows.Scan(&spot.ID, &spot.SpotName, &spot.SpotRegion, &spot.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				failure_res := Response {
					Message: "No spots found",
				}
				json.NewEncoder(w).Encode(failure_res)
				return
			}
			log.Fatal(err)
		}
		spots = append(spots, spot)
	}

	type SpotsResponse struct {
		Message string `json:"message,omitempty"`
		Spots []models.SavedSurfSpot `json:"spots,omitempty"`
	}

	res := SpotsResponse {
		Message: "Success",
		Spots: spots,
	}

	json.NewEncoder(w).Encode(res)
}