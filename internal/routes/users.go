package routes

import (
	"encoding/json"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"project/internal/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	gofakeit.Seed(0)

	users := []models.User{}
	for i := 1; i <= 5; i++ {
		u := models.User{
			ID:    i,
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
