package controller

import (
	"encoding/json"
	"fmt"
	"go-postgres-crud/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []models.ContractWarning `json:"data"`
}

type ResponseSingle struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    models.ContractWarning `json:"data"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ContractWarnings, err := models.AmbilSemuaWarning()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = ContractWarnings

	json.NewEncoder(w).Encode(response)
}

func FindWarning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	period, err := strconv.ParseInt(params["period"], 10, 64)
	if err != nil {
		panic(err)
	}
	day_begin, err := strconv.ParseInt(params["day_begin"], 10, 64)
	if err != nil {
		panic(err)
	}

	colour := params["colour"]

	ContractWarning, err := models.AmbilSatuWarning(period, day_begin, colour)
	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}
	var response ResponseSingle
	response.Status = 1
	response.Message = "OKhaha"
	response.Data = ContractWarning
	json.NewEncoder(w).Encode(response)
}

func InsertWarning(w http.ResponseWriter, r *http.Request) {
	var column models.ContractWarning
	err := json.NewDecoder(r.Body).Decode(&column)
	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}
	dataInsert := models.TambahWarning(column)
	var response ResponseSingle
	response.Status = 1
	response.Message = "OK"
	response.Data = dataInsert
	json.NewEncoder(w).Encode(response)
}

func UpdateWarning(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	period, err := strconv.ParseInt(params["period"], 10, 64)
	if err != nil {
		panic(err)
	}
	day_begin, err := strconv.ParseInt(params["day_begin"], 10, 64)
	if err != nil {
		panic(err)
	}
	colour := params["colour"]

	var column models.ContractWarning

	err = json.NewDecoder(r.Body).Decode(&column)
	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	updateRows, row := models.UpdateWarning(period, day_begin, colour, column)

	var response ResponseSingle
	response.Status = 1
	response.Message = fmt.Sprintf("Berhasil update %v recode", updateRows)
	response.Data = row
	json.NewEncoder(w).Encode(response)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	period, err := strconv.ParseInt(params["period"], 10, 64)
	if err != nil {
		panic(err)
	}
	day_begin, err := strconv.ParseInt(params["day_begin"], 10, 64)
	if err != nil {
		panic(err)
	}
	colour := params["colour"]

	delete := models.DeleteWarning(period, day_begin, colour)

	var response ResponseSingle
	response.Status = 1
	response.Message = fmt.Sprintf("Berhasil hapus %v recode", delete)
	json.NewEncoder(w).Encode(response)
}
