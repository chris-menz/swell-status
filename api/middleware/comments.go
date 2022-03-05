package middleware

import(
	"net/http"
	_ "github.com/lib/pq"
	"api/models"
	"encoding/json"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
)

func GetAllComments(w http.ResponseWriter, r *http.Request){
	db := PostgresConnection()
	defer db.Close()

	var comments []models.Comment = []models.Comment{}

	getCommentsQuery := "SELECT c.* from COMMENTS c"

	rows, err := db.Query(getCommentsQuery)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment

		err = rows.Scan(&comment.ID, &comment.SurfSessionID, &comment.Content, &comment.CommenterID)
		if err == sql.ErrNoRows {
			failure_res := Response {
				Message: "No comments found",
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
		comments = append(comments, comment)
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		log.Fatal(err)
	}
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	var newcomment models.Comment
	var commentid int64

	session, _ := store.Get(r, "session")

	err := json.NewDecoder(r.Body).Decode(&newcomment)
	if err != nil {
		log.Fatal(err)
	}

	db := PostgresConnection()
	defer db.Close()

	insertCommentQuery := `INSERT INTO comments (surf_session_id, content, commenter_id)
		VALUES ($1, $2, $3)
		RETURNING id`

	err = db.QueryRow(insertCommentQuery, newcomment.SurfSessionID, newcomment.Content, session.Values["userID"].(int64)).Scan(&commentid)
	if err != nil {
		log.Fatal(err)
	}

	success_res := Response{
		Message: "Comment added",
	}
	
	err = json.NewEncoder(w).Encode(success_res)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var commenterID int

	session, _ := store.Get(r, "session")
	
	db := PostgresConnection()
	defer db.Close()

	err := db.QueryRow(`SELECT c.commenter_id FROM comments c WHERE id = $1`, id).Scan(&commenterID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		res := Response{
			Message: "can't find comment",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	if int64(commenterID) != session.Values["userID"].(int64) {
		res := Response{
			Message: "can't delete anothers comment",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	
	deleteCommentQuery := `DELETE FROM comments WHERE id=$1`
	sqlresp, err := db.Exec(deleteCommentQuery, id)
	if err != nil {
		log.Fatal(err)
		log.Fatal(sqlresp)
	}

	res := Response {
		Message: "comment " + id + " deleted",	
	}

	json.NewEncoder(w).Encode(res)
}