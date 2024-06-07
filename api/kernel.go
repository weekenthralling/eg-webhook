package api

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"net/http"
	"zjuici.com/tablegpt/eg-webhook/models"
	db "zjuici.com/tablegpt/eg-webhook/storage"
)

func SaveKernel(s *db.KernelSessionStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var session map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		kernel := &models.KernelSession{
			ID:      id,
			Session: datatypes.JSONMap(session),
		}
		if err := s.SaveSession(kernel); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetSession(s *db.KernelSessionStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		kernel, err := s.GetSessionByID(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(kernel)
	}
}

func ListKernels(s *db.KernelSessionStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kernels, err := s.ListSessions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(kernels)
	}
}

func DeleteKernels(s *db.KernelSessionStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ids []string
		if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := s.DeleteSessionsByID(ids); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
