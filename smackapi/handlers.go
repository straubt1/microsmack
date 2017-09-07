package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Category string `json:"Category"`
	Item     string `json:"Item"`
	Value    string `json:"Value"`
}

type Configs []Config

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RUNNING")
}

func returnAllConfigs(w http.ResponseWriter, r *http.Request) {
	configs := Config{Category: "UI", Item: "Background Color", Value: "Blue"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(configs); err != nil {
		panic(err)
	}
}

func returnSingleConfig(w http.ResponseWriter, r *http.Request) {
	// need to implement this...
	vars := mux.Vars(r)
	key := vars["key"]
	var1 := vars["var1"]
	var2 := vars["var2"]

	fmt.Println("Var 1: " + var1)
	fmt.Println("Var 2: " + var2)
	fmt.Fprintf(w, "Key: "+key)
}