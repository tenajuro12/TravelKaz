package controllers

import (
	"diplomaPorject/backend/blogs_service/internal/models"
	"diplomaPorject/backend/blogs_service/utils/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func LikeBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var blog models.Blog
	if err := db.DB.First(&blog, id).Error; err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	blog.Likes++
	db.DB.Save(&blog)
	json.NewEncoder(w).Encode(blog)
}
