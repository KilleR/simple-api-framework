package main

import (
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	api := NewApiResponse(w)

	defer api.Write()

	api.Data["test"] = "Test output data"
}
