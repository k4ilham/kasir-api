package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/model"
	"kasir-api/response"
)

func parseCategoryID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	return strconv.Atoi(idStr)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseCategoryID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Category Id")
		return
	}
	for _, c := range model.Categories {
		if c.ID == id {
			response.WriteJSON(w, http.StatusOK, c)
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Category tidak ditemukan")
}

func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseCategoryID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Category Id")
		return
	}
	var update model.Category
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	for i := range model.Categories {
		if model.Categories[i].ID == id {
			update.ID = id
			model.Categories[i] = update
			response.WriteJSON(w, http.StatusOK, update)
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Category tidak ditemukan")
}

func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseCategoryID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid Category Id")
		return
	}
	for i, c := range model.Categories {
		if c.ID == id {
			model.Categories = append(model.Categories[:i], model.Categories[i+1:]...)
			response.WriteJSON(w, http.StatusOK, map[string]string{
				"message": "sukses delete category",
				"status":  "Ok",
			})
			return
		}
	}
	response.WriteError(w, http.StatusNotFound, "Category tidak ditemukan")
}

func CategoryItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetCategoryByID(w, r)
	case http.MethodPut:
		UpdateCategoryByID(w, r)
	case http.MethodDelete:
		DeleteCategoryByID(w, r)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func CategoryCollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response.WriteJSON(w, http.StatusOK, model.Categories)
	case http.MethodPost:
		var c model.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			response.WriteError(w, http.StatusBadRequest, "Invalid Request")
			return
		}
		c.ID = len(model.Categories) + 1
		model.Categories = append(model.Categories, c)
		response.WriteJSON(w, http.StatusCreated, c)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
