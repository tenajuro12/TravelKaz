package controllers

import (
	"diplomaPorject/backend/blogs_service/internal/models"
	"diplomaPorject/backend/blogs_service/utils/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	comment.BlogID = uint(id)
	db.DB.Create(&comment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}
