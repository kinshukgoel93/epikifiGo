package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthCheckResponse struct {
	Msg  string
	Code int
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := healthCheckResponse{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}
