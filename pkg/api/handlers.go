package api

import (
	"encoding/json"
	"net/http"

	"api-articles/pkg/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ArticlesRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/articles", GetArticles(db)).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticle(db)).Methods("GET")
	router.HandleFunc("/articles", CreateArticle(db)).Methods("POST")
	router.HandleFunc("/articles/{id}", UpdateArticle(db)).Methods("PUT")
	router.HandleFunc("/articles/{id}", DeleteArticle(db)).Methods("DELETE")
}

// GetArticles fetches all articles
func GetArticles(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var articles []models.Article
		if err := db.Find(&articles).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// gimme example reponse
		response := map[string]interface{}{
			"status": "success",
			"data":   articles,
		}

		json.NewEncoder(w).Encode(response)
	}
}

// GetArticle fetches a article by ID
func GetArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var article models.Article
		if err := db.First(&article, id).Error; err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}

		// gimme example reponse
		response := map[string]interface{}{
			"status": "success",
			"data":   article,
		}

		json.NewEncoder(w).Encode(response)
	}
}

// CreateArticle creates a new article
func CreateArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article models.Article
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&article).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// gimme example reponse
		response := map[string]interface{}{
			"status": "success created article",
			"data":   article,
		}

		json.NewEncoder(w).Encode(response)
	}
}

// UpdateArticle updates an existing article
func UpdateArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var article models.Article
		if err := db.First(&article, id).Error; err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Save(&article).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// gimme example reponse
		response := map[string]interface{}{
			"status": "success updated article",
			"data":   article,
		}

		json.NewEncoder(w).Encode(response)
	}
}

// DeleteArticle deletes a article
func DeleteArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := db.Delete(&models.Article{}, id).Error; err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}

		// gimme example reponse
		response := map[string]interface{}{
			"status": "success deleted article",
		}

		json.NewEncoder(w).Encode(response)
	}
}
