package models

//models to insert data to postgres

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type SurfSession struct {
	ID int64 `json:"id"`
	Description string `json:"description"`
	DateTime string `json:"datetime"`
	Swell string `json:"swell"`
	Wind string `json:"wind"`
	Tide string `json:"tide"`
	SurfSpot string `json:"surf_spot"`
	SpotLocation string `json:"spot_location"`
	SpotRegion string `json:"spot_region"`
	CreatorID int64 `json:"creator_id"`
	SurfboardID *int64 `json:"surfboard_id"`
	IsPublic bool `json:"is_public"`
}

type Favorite struct {
	ID int64 `json:"id"`
	SurfSessionID int64 `json:"surf_session_id"`
	UserID int64 `json:"user_id"`
}

type Comment struct {
	ID int64 `json:"id"`
	SurfSessionID int64 `json:"surf_session_id"`
	Content string `json:"content"`
	CommenterID int64 `json:"commenter_id"`
}

type SavedSurfSpot struct {
	ID int64 `json:"id"`
	SpotName string `json:"spot_name"`
	SpotRegion string `json:"spot_region"`
	UserID int64 `json:"user_id"`
}

type Surfboard struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	HeightInches string `json:"height_inches"`
	WidthInches string `json:"width_inches"`
	ThicknessInches string `json:"thickness_inches"`
	ConstructionMaterial string `json:"construction_material"`
	UserID int64 `json:"user_id"`
}

//response models to 'get' requests from client

type UserResponse struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

//auto generated using json-to-go https://mholt.github.io/json-to-go/

type SgResponse struct {
	Hours []struct {
		SecondarySwellDirection struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"secondarySwellDirection"`
		SecondarySwellHeight struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"secondarySwellHeight"`
		SecondarySwellPeriod struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"secondarySwellPeriod"`
		SwellDirection struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellDirection"`
		SwellHeight struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellHeight"`
		SwellPeriod struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellPeriod"`
		Time string `json:"time"`
		WaterTemperature struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waterTemperature"`
		WaveDirection struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waveDirection"`
		WaveHeight struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waveHeight"`
		WavePeriod struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"wavePeriod"`
		WindDirection struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"windDirection"`
		WindSpeed struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"windSpeed"`
		WindWaveDirection struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"windWaveDirection"`
		WindWaveHeight struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"windWaveHeight"`
		WindWavePeriod struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"windWavePeriod"`
	} `json:"hours"`
	Meta struct {
		Cost         int      `json:"cost"`
		DailyQuota   int      `json:"dailyQuota"`
		End          string   `json:"end"`
		Lat          float64  `json:"lat"`
		Lng          float64  `json:"lng"`
		Params       []string `json:"params"`
		RequestCount int      `json:"requestCount"`
		Source       []string `json:"source"`
		Start        string   `json:"start"`
	} `json:"meta"`
}

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64     `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}



type OWMResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Hourly []struct {
		Dt         int     `json:"dt"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Pop int `json:"pop"`
	} `json:"hourly"`
	Daily []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    int     `json:"pop"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
}

type TimeZoneResponse struct {
	Status           string `json:"status"`
	Message          string `json:"message"`
	CountryCode      string `json:"countryCode"`
	CountryName      string `json:"countryName"`
	RegionName       string `json:"regionName"`
	CityName         string `json:"cityName"`
	ZoneName         string `json:"zoneName"`
	Abbreviation     string `json:"abbreviation"`
	GmtOffset        int    `json:"gmtOffset"`
	Dst              string `json:"dst"`
	ZoneStart        int    `json:"zoneStart"`
	ZoneEnd          int    `json:"zoneEnd"`
	NextAbbreviation string `json:"nextAbbreviation"`
	Timestamp        int    `json:"timestamp"`
	Formatted        string `json:"formatted"`
}

type TideResponse struct {
	Status        int     `json:"status"`
	CallCount     int     `json:"callCount"`
	Copyright     string  `json:"copyright"`
	RequestLat    float64 `json:"requestLat"`
	RequestLon    float64 `json:"requestLon"`
	ResponseLat   float64 `json:"responseLat"`
	ResponseLon   float64 `json:"responseLon"`
	Atlas         string  `json:"atlas"`
	Station       string  `json:"station"`
	Timezone      string  `json:"timezone"`
	RequestDatum  string  `json:"requestDatum"`
	ResponseDatum string  `json:"responseDatum"`
	Extremes      []struct {
		Dt     int     `json:"dt"`
		Date   string  `json:"date"`
		Height float64 `json:"height"`
		Type   string  `json:"type"`
	} `json:"extremes"`
}

