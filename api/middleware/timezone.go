package middleware

import(
	"net/http"
	"github.com/gorilla/mux"
	"api/models"
	"encoding/json"
	"io"
	"os"
	"log"
)

func GetTimeZoneOffset(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])

	var timeZoneResponse models.TimeZoneResponse

	response, err := http.Get("http://vip.timezonedb.com/v2.1/get-time-zone?key=" + os.Getenv("TIMEZONE_KEY") + "&format=json&by=position&lat=" + lat + "&lng=" + lng)

	if err != nil {
        log.Fatal(err)
    }

	defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	err = json.Unmarshal(responseData, &timeZoneResponse)
	if err != nil {
        log.Fatal(err)
    }

	err = json.NewEncoder(w).Encode(timeZoneResponse)
	if err != nil {
		log.Fatal(err)
	}
}