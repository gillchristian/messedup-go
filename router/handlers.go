package main

import (
	"encoding/json"
	"gateway/services"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index is the home route
func Index(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Wellcome to the API",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}

// ListServices gets all the services
func ListServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services.services)
}

// FindService by name
func FindService(w http.ResponseWriter, r *http.Request) {
	routeParams := mux.Vars(r)
	if name, err := strconv.Atoi(routeParams["service"]); err != nil {
		panic(err)
	}
	service := services.Find(name)
	if len(service.name) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewDecoder(r).Encode(service); err != nil {
			panic(err)
		}
		return
	}

	// 404 if not found
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	response := map[string]string{
		"status":  string(http.StatusNotFound),
		"message": "Service Not found",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// RegisterService to be routed
func RegisterService(w http.ResponseWriter, r *http.Request) {
	var service Service
	if body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
		panic(err)
	}
	if err := body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &service); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		json.NewEncoder(w).Encode(err)
	}

	t := services.Add(service)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// RemoveService to be routed
func RemoveService(w http.ResponseWriter, r *http.Request) {
	var service Service
	if body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
		panic(err)
	}
	if err := body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &service); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		json.NewEncoder(w).Encode(err)
	}

	t := services.Add(service)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
