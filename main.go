package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

const (
	PORT = ":4444"
	MIN  = 1
	MAX  = 100
)

type WinterData struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/winter", getWinter)
	http.ListenAndServe(PORT, nil)
}

func getWinter(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		water := rand.Intn(MAX-MIN) + MIN
		wind := rand.Intn(MAX-MIN) + MIN
		winterStatus := "DEFAULT"

		if water < 5 {
			winterStatus = "AMAN"
		} else if water >= 6 && water == 8 {
			winterStatus = "SIAGA"
		} else if water > 8 {
			winterStatus = "BAHAYA"
		} else if wind < 6 {
			winterStatus = "AMAN"
		} else if wind >= 7 && wind == 15 {
			winterStatus = "SIAGA"
		} else if wind > 15 {
			winterStatus = "BAHAYA"
		}

		result := WinterData{
			Water:  water,
			Wind:   wind,
			Status: winterStatus,
		}

		tmpl, err := template.ParseFiles("views/winter_status.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, result)
		return

	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
