package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	// "strconv"
	// "time"
)

type Weather struct {
}

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    WeatherData `json:"data"`
}

type WeatherData struct {
	Forecast24h map[int]WeatherObject `json:"forecast_24h"`
}

type WeatherObject struct {
	DayWeather       string `json:"day_weather"`
	DayWindDirection string `json:"day_wind_direction"`
	Time             string `json:"time"`
}

func (w *Weather) GetData() (string, error) {
	s, _ := time.LoadLocation("Asia/Shanghai")
	d, _ := base64.StdEncoding.DecodeString(weatherUrl)
	str := string(d) + weatherPath
	param := url.Values{}
	u, _ := url.Parse(str)
	param.Set("source", "pc")
	param.Set("weather_type", "forecast_1h|forecast_24h|index|alarm|limit|tips|rise")
	param.Set("province", "北京市")
	param.Set("city", "北京市")
	u.RawQuery = param.Encode()
	uPath := u.String()
	resp, err := http.Get(uPath)
	data, _ := ioutil.ReadAll(resp.Body)
	info := Result{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Println(err.Error())
	}
	z := ""
	for i := 1; i < len(info.Data.Forecast24h); i++ {
		now, _ := time.ParseInLocation("2006-01-02", info.Data.Forecast24h[i].Time, s)
		z += fmt.Sprintf(`
<table><tr><tr><td width='80'>%s</td><td width='80'>%s</td><td width='50'>%s</td><td width='50'>%s</td></tr></tr></table>
`, now.Weekday(), info.Data.Forecast24h[i].Time, info.Data.Forecast24h[i].DayWeather, info.Data.Forecast24h[i].DayWindDirection)
	}
	return z, err
}
