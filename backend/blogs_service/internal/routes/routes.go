package routes

import (
	"diplomaPorject/backend/blogs_service/internal/controllers"
	"github.com/gorilla/mux"
)

// Register all blog routes
func RegisterBlogRoutes(r *mux.Router) {
	r.HandleFunc("/blogs", controllers.GetBlogs).Methods("GET")
	r.HandleFunc("/blogs/{id:[0-9]+}", controllers.GetBlog).Methods("GET")
	r.HandleFunc("/blogs", controllers.CreateBlog).Methods("POST")
	r.HandleFunc("/blogs/{id:[0-9]+}", controllers.UpdateBlog).Methods("PUT")
	r.HandleFunc("/blogs/{id:[0-9]+}", controllers.DeleteBlog).Methods("DELETE")
	r.HandleFunc("/blogs/{id:[0-9]+}/like", controllers.LikeBlog).Methods("POST")
	r.HandleFunc("/blogs/{id:[0-9]+}/comment", controllers.AddComment).Methods("POST")
}
