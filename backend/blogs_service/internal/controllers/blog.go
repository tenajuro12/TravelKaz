package controllers

import (
	"diplomaPorject/backend/blogs_service/internal/models"
	"diplomaPorject/backend/blogs_service/utils/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db.DB.Create(&blog)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blog)
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs models.Blog
	db.DB.Preload("Comments").Find(&blogs)
	json.NewEncoder(w).Encode(blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var blog models.Blog

	if err := db.DB.Preload("Comments").First(&blog, id).Error; err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var blog models.Blog
	if err := db.DB.First(&blog, id).Error; err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db.DB.Save(&blog)
	json.NewEncoder(w).Encode(blog)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := db.DB.Delete(&models.Blog{}, id).Error; err != nil {
		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
