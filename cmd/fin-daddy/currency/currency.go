package currency

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const url = "https://www.cbr-xml-daily.ru/daily_json.js"

func GetCurrency(code string) (string, float64) {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		log.Print(err)
		return "", 0
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Print(err)
		return "", 0
	}

	var j responseScheme
	err = json.Unmarshal(b, &j)
	if err != nil {
		log.Print(err)
		return "", 0
	}

	if v, ok := j.Valute[strings.ToUpper(code)]; ok {
		return v.Name, v.Value
	}

	return "", 0
}

type responseScheme struct {
	Valute map[string]curr `json:"Valute"`
}

type curr struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}
