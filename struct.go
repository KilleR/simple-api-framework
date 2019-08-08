package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiResponse struct {
	w      http.ResponseWriter
	r      *http.Request
	Status int
	Data   map[string]interface{}
	Error  string
}

func NewApiResponse(w http.ResponseWriter) (api *ApiResponse) {
	api = new(ApiResponse)
	api.Init(w)
	return
}

func (api *ApiResponse) Init(w http.ResponseWriter) {
	api.w = w
	api.Data = make(map[string]interface{})
}

func (api ApiResponse) Write() {
	if api.Status != 0 {
		api.w.WriteHeader(api.Status)
	} else {
		api.w.WriteHeader(http.StatusOK)
	}

	outData, err := json.Marshal(api)
	if err != nil {
		api.w.WriteHeader(500)
		log.Println("Failed to marshal response:", err)
	}

	_, err = api.w.Write(outData)
	if err != nil {
		log.Println("Failed to write API output", err)
	}
}

func (api *ApiResponse) Fail(errorText string, status int) {

}
