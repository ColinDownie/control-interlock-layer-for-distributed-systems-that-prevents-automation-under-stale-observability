package main

import (
	"encoding/json"
	"net/http"
)

type SRM struct {
	Omega   float64 `json:"omega"`
	TauObs  float64 `json:"tau_obs"`
	TauAct  float64 `json:"tau_act"`
	G       float64 `json:"g"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	s := SRM{
		Omega:  0.8,
		TauObs: 0.4,
		TauAct: 0.3,
	}

	s.G = (s.TauObs + s.TauAct) * s.Omega

	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/signal", handler)
	http.ListenAndServe(":8080", nil)
}
