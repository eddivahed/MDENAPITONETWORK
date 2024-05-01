package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/network-api/fabric"
)

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = fabric.RegisterUser(user.Username, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}