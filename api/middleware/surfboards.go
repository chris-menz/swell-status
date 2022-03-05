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

func CreateSurfboard(w http.ResponseWriter, r *http.Request) {
	var newsurfboard models.Surfboard
	var surfboardid int64

	session, _ := store.Get(r, "session")

	err := json.NewDecoder(r.Body).Decode(&newsurfboard)
	if err != nil {
		log.Fatal(err)
	}

	if newsurfboard.Name == "" {
		res := Response {
			Message: "Can't add surfboard without name",
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	db := PostgresConnection()
	defer db.Close()

	insertSurfBoardQuery := `INSERT INTO surfboards (name, height_inches, width_inches, thickness_inches, construction_material, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err = db.QueryRow(insertSurfBoardQuery, newsurfboard.Name, newsurfboard.HeightInches, newsurfboard.WidthInches, newsurfboard.ThicknessInches, newsurfboard.ConstructionMaterial, session.Values["userID"].(int64)).Scan(&surfboardid)
	if err != nil {
		log.Fatal(err)
	}

	success_res := Response{
		Message: "Surfboard added",
	}
	
	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSurfboardByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var surfboard models.Surfboard

	db := PostgresConnection()
	defer db.Close()

	getSurfboardQuery := `SELECT * FROM surfboards WHERE id = $1`

	row := db.QueryRow(getSurfboardQuery, id)
	err := row.Scan(&surfboard.ID, &surfboard.Name, &surfboard.HeightInches, &surfboard.WidthInches, &surfboard.ThicknessInches, &surfboard.ConstructionMaterial, &surfboard.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			failure_res := Response {
				Message: "No surfboard found",
			}
			json.NewEncoder(w).Encode(failure_res)
			return
		}
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(surfboard)
}

func GetSurfboardsByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id := params["user_id"]
	var surfboards []models.Surfboard

	db := PostgresConnection()
	defer db.Close()

	getSurfboardsQuery := `SELECT s.id, s.name, s.height_inches, s.width_inches, s.thickness_inches, s.construction_material, s.user_id
		FROM surfboards s INNER JOIN users u ON s.user_id = u.id 
		WHERE s.user_id = $1`

	rows, err := db.Query(getSurfboardsQuery, user_id)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var surfboard models.Surfboard

		err = rows.Scan(&surfboard.ID, &surfboard.Name, &surfboard.HeightInches, &surfboard.WidthInches, &surfboard.ThicknessInches, &surfboard.ConstructionMaterial, &surfboard.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				failure_res := Response {
					Message: "No surfboards found",
				}
				json.NewEncoder(w).Encode(failure_res)
				return
			}
			log.Fatal(err)
		}
		surfboards = append(surfboards, surfboard)
	}

	json.NewEncoder(w).Encode(surfboards)
}

func EditSurfboard(w http.ResponseWriter, r *http.Request) {
	db := PostgresConnection()
	defer db.Close()
	
	params := mux.Vars(r)
	id := params["id"]
	var surfboard models.Surfboard
	var surfboardCreatorID int

	session, _ := store.Get(r, "session")

	err := db.QueryRow(`SELECT surfboards.user_id FROM surfboards WHERE id = $1`, id).Scan(&surfboardCreatorID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find surfboard",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&surfboard)
	if err != nil {
		log.Fatal(err)
	}
	
	if surfboardCreatorID != session.Values["userID"].(int) {
		res := Response{
			Message: "can't edit anothers surfboard",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	
	updateSurfboardQuery := `UPDATE surfboards SET name=$2, height_inches=$3, 
	width_inches=$4, thickness_inches=$5, construction_material=$6, user_id=$7
	WHERE id=$1`

	sqlresp, err := db.Exec(updateSurfboardQuery, id, surfboard.Name, surfboard.HeightInches, surfboard.WidthInches, surfboard.ThicknessInches, surfboard.ConstructionMaterial, session.Values["userID"].(int))
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "board updated",	
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteSurfboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var surfboardCreatorID int

	session, _ := store.Get(r, "session")
	
	db := PostgresConnection()
	defer db.Close()

	err := db.QueryRow(`SELECT surfboards.user_id FROM surfboards WHERE id = $1`, id).Scan(&surfboardCreatorID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find surfboard",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	if int64(surfboardCreatorID) != session.Values["userID"].(int64) {
		res := Response{
			Message: "can't delete anothers surfboard",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	deleteSurfboardQuery := `DELETE FROM surfboards WHERE id=$1`
	sqlresp, err := db.Exec(deleteSurfboardQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "board " + id + " deleted",	
	}

	json.NewEncoder(w).Encode(res)
}