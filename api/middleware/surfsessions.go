package middleware

import (
	"net/http"
	"api/models"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"database/sql"
)

type SurfSessionsResponse struct {
	Message string `json:"message,omitempty"`
	SurfSessions []models.SurfSession `json:"surf_sessions,omitempty"`
}


func CreateSurfSession(w http.ResponseWriter, r *http.Request) {
	var newSession models.SurfSession
	var newSessionID int64


	session, _ := store.Get(r, "session")

	err := json.NewDecoder(r.Body).Decode(&newSession)
	if err != nil {
		log.Fatal(err)
	}

	missingParameters := 
		newSession.Description == "" || 
		newSession.DateTime == "" ||
		newSession.Swell == "" ||
		newSession.Wind == "" ||
		newSession.Tide == "" ||
		newSession.SurfSpot == "" ||
		newSession.SpotLocation == "" ||
		newSession.SpotRegion == ""

	if missingParameters {
		res := Response {
			Message: "Missing info to create surf session",
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	db := PostgresConnection()
	defer db.Close()

	insertSurfSessionQuery := `INSERT INTO surf_sessions 
		(description, datetime, swell, wind, tide, surf_spot, spot_region, spot_location, creator_id, surfboard_id, is_public)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`
	
	err = db.QueryRow(insertSurfSessionQuery, 
		newSession.Description, newSession.DateTime, newSession.Swell,
		newSession.Wind, newSession.Tide, newSession.SurfSpot, newSession.SpotRegion,
		newSession.SpotLocation, session.Values["userID"].(int64),
		newSession.SurfboardID, newSession.IsPublic).Scan(&newSessionID)
	if err != nil {
		log.Fatal(err)
	}

	success_res := Response {
		Message: "Session added",
	}
	
	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSurfSessionsByUserId(w http.ResponseWriter, r *http.Request) {
	var surfSessions []models.SurfSession
	params := mux.Vars(r)
	user_id := params["user_id"]

	db := PostgresConnection()
	defer db.Close()

	getSurfSessionsQuery := `SELECT s.id, s.description, s.datetime, s.swell, s.wind, s.tide, 
	s.surf_spot, s.spot_region, s.creator_id, s.surfboard_id, s.is_public
	FROM surf_sessions s INNER JOIN users u ON s.creator_id = u.id 
	WHERE s.creator_id = $1`

	rows, err := db.Query(getSurfSessionsQuery, user_id)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var surfSession models.SurfSession

		err = rows.Scan(&surfSession.ID, &surfSession.Description, &surfSession.DateTime, &surfSession.Swell, 
			&surfSession.Wind, &surfSession.Tide, &surfSession.SurfSpot, &surfSession.SpotRegion, &surfSession.CreatorID, 
			&surfSession.SurfboardID, &surfSession.IsPublic)
		if err != nil {
			if err == sql.ErrNoRows {
				failure_res := Response {
					Message: "No surf sessions found",
				}
				json.NewEncoder(w).Encode(failure_res)
				return
			}
			log.Fatal(err)
		}
		surfSessions = append(surfSessions, surfSession)
	}

	res := SurfSessionsResponse {
		Message: "success",
		SurfSessions: surfSessions, 
	}

	json.NewEncoder(w).Encode(res)
}

func GetPublicSurfSessions(w http.ResponseWriter, r *http.Request) {
	var surfSessions []models.SurfSession = make([]models.SurfSession, 0)

	db := PostgresConnection()
	defer db.Close()

	getSurfSessionsQuery := `SELECT s.*
	FROM surf_sessions s`

	rows, err := db.Query(getSurfSessionsQuery)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var surfSession models.SurfSession

		err = rows.Scan(&surfSession.ID, &surfSession.Description, &surfSession.DateTime, &surfSession.Swell, 
			&surfSession.Wind, &surfSession.Tide, &surfSession.SurfSpot, &surfSession.SpotRegion, &surfSession.SpotLocation, &surfSession.CreatorID, 
			&surfSession.SurfboardID, &surfSession.IsPublic)
		if err != nil {
			log.Fatal(err)
		}
		surfSessions = append(surfSessions, surfSession)
	}

	res := SurfSessionsResponse {
		Message: "success",
		SurfSessions: surfSessions, 
	}

	if len(surfSessions) == 0 {
		res = SurfSessionsResponse {
			Message: "no sessions found",
		}
	}

	json.NewEncoder(w).Encode(res)
}

func RemoveSurfSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var surfSessionCreatorID int

	session, _ := store.Get(r, "session")
	
	db := PostgresConnection()
	defer db.Close()

	err := db.QueryRow(`SELECT s.creator_id FROM surf_sessions s WHERE id = $1`, id).Scan(&surfSessionCreatorID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find surf session",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if int64(surfSessionCreatorID) != session.Values["userID"].(int64) {
		res := Response{
			Message: "can't delete anothers surf session",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	deleteFavoritesQuery := `DELETE FROM favorites WHERE surf_session_id=$1`
	sqlresp, err := db.Exec(deleteFavoritesQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	deleteCommentsQuery := `DELETE FROM comments WHERE surf_session_id=$1`
	sqlresp, err = db.Exec(deleteCommentsQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	deleteSurfSessionQuery := `DELETE FROM surf_sessions WHERE id=$1`
	sqlresp, err = db.Exec(deleteSurfSessionQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "session deleted",	
	}

	json.NewEncoder(w).Encode(res)
}