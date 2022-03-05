package middleware

import (
	"net/http"
	"io"
	"os"
	"log"
	"encoding/json"
	"api/models"
	"time"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
)

//Stormglass api
func getSgData(url string) models.SgResponse {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", os.Getenv("SG_KEY"))
	response, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

	defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	var sgResponse models.SgResponse
	err = json.Unmarshal(responseData, &sgResponse)
	if err != nil {
        log.Fatal(err)
    }

    return sgResponse
}

//Open Weather Map api current weather
func getWeatherData(url string) models.WeatherResponse {
	response, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

	defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	var weatherResponse models.WeatherResponse
	err = json.Unmarshal(responseData, &weatherResponse)
	if err != nil {
        log.Fatal(err)
    }

    return weatherResponse
}

// Open Weather Map api 7 day forecast 
func getOWMData(url string) models.OWMResponse {
	response, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

	defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	var owmResponse models.OWMResponse
	err = json.Unmarshal(responseData, &owmResponse)
	if err != nil {
        log.Fatal(err)
    }

    return owmResponse
}

func getTidesData(url string) models.TideResponse{
	response, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

	defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	var tideResponse models.TideResponse
	err = json.Unmarshal(responseData, &tideResponse)
	if err != nil {
        log.Fatal(err)
    }

    return tideResponse
}


// handlers

func SgHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])

	cache := Cache()
	defer cache.Close()

	var sg models.SgResponse

	key := "sg" + lat + "," + lng

	_, err := cache.Get(key).Result()

	if err == redis.Nil {
		sg = getSgData("https://api.stormglass.io/v2/weather/point?lat=" + lat + "&lng=" + lng + "&params=swellHeight,swellPeriod,swellDirection,windSpeed,windDirection,waterTemperature,waveHeight,wavePeriod,waveDirection,windWaveHeight,windWavePeriod,windWaveDirection,secondarySwellHeight,secondarySwellPeriod,secondarySwellDirection&source=noaa,sg")
		jsondata, _ := json.Marshal(sg)

		// key will last until the end of the utc day. This api updates everday at utc 0:00, 
		// so no need to re fetch data from it until new day
		cacheTime := (time.Duration(24 - time.Now().UTC().Hour()) * time.Hour + 60) - (time.Duration(time.Now().UTC().Minute()) * time.Minute)
		
		err := cache.Set(key, jsondata, cacheTime).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		jsonstring, err := cache.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		
		json.Unmarshal([]byte(jsonstring), &sg)
	}
	err = json.NewEncoder(w).Encode(sg)
	if err != nil {
		log.Fatal(err)
	}
}

func PastSgHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])
	start := string(params["start"])
	end := string(params["end"])

	cache := Cache()
	defer cache.Close()

	var sg models.SgResponse

	key := "pastSg" + lat + "," + lng + "," + start + "," + end

	_, err := cache.Get(key).Result()

	if err == redis.Nil {
		sg = getSgData("https://api.stormglass.io/v2/weather/point?lat=" + lat + "&lng=" + lng + "&start=" + start + "&end=" + end + "&params=swellHeight,swellPeriod,swellDirection,windSpeed,windDirection,waterTemperature,waveHeight,wavePeriod,waveDirection,windWaveHeight,windWavePeriod,windWaveDirection,secondarySwellHeight,secondarySwellPeriod,secondarySwellDirection&source=noaa,sg")
		jsondata, _ := json.Marshal(sg)

		cacheTime := (time.Duration(7 * 24 * time.Hour))
		
		err := cache.Set(key, jsondata, cacheTime).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		jsonstring, err := cache.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		
		json.Unmarshal([]byte(jsonstring), &sg)
	}
	err = json.NewEncoder(w).Encode(sg)
	if err != nil {
		log.Fatal(err)
	}
}

func WeatherHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])

	cache := Cache()
	defer cache.Close()

	var weather models.WeatherResponse

	key := "owm" + lat + "," + lng

	_, err := cache.Get(key).Result()
	
	if err == redis.Nil {
		weather = getWeatherData("https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lng + "&appid=" + os.Getenv("OWM_KEY") + "&units=imperial")
		jsondata, _ := json.Marshal(weather)

		//only cache for 1 hour due to volatility of weather
		err := cache.Set(key, jsondata, time.Hour).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		jsonstring, err := cache.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		
		json.Unmarshal([]byte(jsonstring), &weather)
	}
	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		log.Fatal(err)
	}
}

func OWMHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])

	cache := Cache()
	defer cache.Close()

	var owm models.OWMResponse

	key := "owm" + lat + "," + lng

	_, err := cache.Get(key).Result()
	
	if err == redis.Nil {
		owm = getOWMData("https://api.openweathermap.org/data/2.5/onecall?lat=" + lat + "&lon=" + lng + "&appid=" + os.Getenv("OWM_KEY") + "&units=imperial&exclude=minutely,alerts")
		jsondata, _ := json.Marshal(owm)

		//only cache for 1 hour due to volatility of weather
		err := cache.Set(key, jsondata, time.Hour).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		jsonstring, err := cache.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		
		json.Unmarshal([]byte(jsonstring), &owm)
	}
	err = json.NewEncoder(w).Encode(owm)
	if err != nil {
		log.Fatal(err)
	}
}

func TideHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lat := string(params["lat"])
	lng := string(params["lng"])
	date := string(params["date"])

	cache := Cache()
	defer cache.Close()

	var tides models.TideResponse

	key := "tides" + lat + "," + lng + "," + date

	_, err := cache.Get(key).Result()

	if err == redis.Nil {
		tides = getTidesData("https://www.worldtides.info/api/v3?extremes&lat=" + lat + "&lon=" + lng + "&key=" +os.Getenv("TIDE_KEY") + "&days=10&localtime&datum=CD&date=" + date)
		jsondata, _ := json.Marshal(tides)

		cacheTime := (time.Duration(24 - time.Now().UTC().Hour()) * time.Hour + 60) - (time.Duration(time.Now().UTC().Minute()) * time.Minute)
		
		err := cache.Set(key, jsondata, cacheTime).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		jsonstring, err := cache.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		
		json.Unmarshal([]byte(jsonstring), &tides)
	}
	err = json.NewEncoder(w).Encode(tides)
	if err != nil {
		log.Fatal(err)
	}
}


