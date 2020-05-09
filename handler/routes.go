package handler

import (
	"net/http"
	"todo/db"
)

func SetUpRouting() *http.ServeMux {
	todoHandler := &todoHandler{
		samples: &db.Sample{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/samples", todoHandler.GetSamples)

	return mux
}
