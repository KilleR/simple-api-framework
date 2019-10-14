package main

import (
	simpleApiFramework "github.com/KilleR/simple-api-framework"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	api := simpleApiFramework.NewApiResponse(w)

	defer api.Write()

	api.Data["test"] = "Test output data"
}
