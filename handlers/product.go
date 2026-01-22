package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/model"
	"kasir-api/response"
)

func parseProdukID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	return strconv.Atoi(idStr)
}

func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseProdukID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Produk Id")
		return
	}
	for _, p := range model.Products {
		if p.ID == id {
			response.WriteJSON(w, http.StatusOK, p)
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Produk belum ada")
}

func UpdateProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseProdukID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Produk Id")
		return
	}
	var updateProduk model.Produk
	if err := json.NewDecoder(r.Body).Decode(&updateProduk); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if !model.CategoryExists(updateProduk.CategoryID) {
		response.WriteError(w, http.StatusBadRequest, "Category tidak ditemukan")
		return
	}
	for i := range model.Products {
		if model.Products[i].ID == id {
			updateProduk.ID = id
			model.Products[i] = updateProduk
			response.WriteJSON(w, http.StatusOK, updateProduk)
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Produk belum ada")
}

func DeleteProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseProdukID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Produk Id")
		return
	}
	for i, p := range model.Products {
		if p.ID == id {
			model.Products = append(model.Products[:i], model.Products[i+1:]...)
			response.WriteJSON(w, http.StatusOK, map[string]string{
				"message": "sukses delete produk",
				"status":  "Ok",
			})
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Produk belum ada")
}

func ProdukItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProdukByID(w, r)
	case http.MethodPut:
		UpdateProdukByID(w, r)
	case http.MethodDelete:
		DeleteProdukByID(w, r)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func ProdukCollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response.WriteJSON(w, http.StatusOK, model.Products)
	case http.MethodPost:
		var produkBaru model.Produk
		if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
			response.WriteError(w, http.StatusBadRequest, "Invalid Request")
			return
		}
		if !model.CategoryExists(produkBaru.CategoryID) {
			response.WriteError(w, http.StatusBadRequest, "Category tidak ditemukan")
			return
		}
		produkBaru.ID = len(model.Products) + 1
		model.Products = append(model.Products, produkBaru)
		response.WriteJSON(w, http.StatusCreated, produkBaru)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
