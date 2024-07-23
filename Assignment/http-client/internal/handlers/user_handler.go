package handlers

import (
	"database/Assignment/http-client/internal/models"
	"database/Assignment/http-client/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo repository.UserRepository
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// GetUsers handles GET requests to /users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling user request...") // Basic logging
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

// GetUser handles GET requests to /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userRepo.GetUserByID(id)
	if err != nil {
		if err == repository.ErrUserNotFound {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// CreateUser handles POST requests to /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := h.userRepo.CreateUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// UpdateUser handles PUT requests to /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user.ID = id

	if err := h.userRepo.UpdateUser(&user); err != nil {
		if err == repository.ErrUserNotFound {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User updated"})
}

// DeleteUser handles DELETE requests to /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.userRepo.DeleteUser(id); err != nil {
		if err == repository.ErrUserNotFound {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
}

func (h *UserHandler) SearchUsersByAddress(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling SearchUsersByAddress request...") // Basic logging

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	// Get pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1 // Default to page 1 if not provided
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 10 // Default to page size 10 if not provided
	}

	users, err := h.userRepo.SearchUsersByAddress(address, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}
