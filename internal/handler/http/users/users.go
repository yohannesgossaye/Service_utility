package users

import (
	"encoding/json"
	"net/http"

	"users/internal/domain/dto"
	"users/internal/handler/http/users/core"
	"users/internal/service"

	"github.com/go-chi/chi/v5"
)

type UserH struct {
	userservice service.UserS
}

func NewUserH(us service.UserS) *UserH {
	return &UserH{
		userservice: us,
	}
}

func (h *UserH) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UsercreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newCreate := core.Tomodelreq(req)
	create, err := h.userservice.CreateUser(r.Context(), newCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := dto.UserResponse{
		ID:             create.ID,
		FullName:       create.FullName,
		Email:          create.Email,
		Account_number: create.Account_number,
		Phone_number:   create.Phone_number,
		Balance:        create.Balance,
		Is_active:      create.Is_active,
	}
	json.NewEncoder(w).Encode(struct {
		Status  int              `json:"status"`
		Message string           `json:"message"`
		Data    dto.UserResponse `json:"data"`
	}{
		Status:  http.StatusCreated,
		Message: "user created successfully",
		Data:    response,
	})
}

func (h *UserH) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userservice.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (h *UserH) GetUser(w http.ResponseWriter, r *http.Request) {
	idparam := chi.URLParam(r, "id")

	user, err := h.userservice.GetUser(r.Context(), idparam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: "user fetched successfully",
		Data:    user,
	})
}

func (h *UserH) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idparam := chi.URLParam(r, "id")
	var req dto.UsercreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	NewUpdate := core.Tomodelreq(req)
	updated, err := h.userservice.UpdateUsers(r.Context(), idparam, &NewUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: "user updated successfully",
		Data:    updated,
	})
}

func (h *UserH) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idparam := chi.URLParam(r, "id")

	_, err := h.userservice.DeleteUser(r.Context(), idparam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "user deleted successfully",
	})
}

func (h *UserH) EnableAccount(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	_, err := h.userservice.EnableAccount(r.Context(), idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "user enabled successfully",
	})
}
