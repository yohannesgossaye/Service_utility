package users

import (
	"encoding/json"
	"net/http"

	"users/internal/domain/dto"
	"users/internal/handler/http/users/core"
	"users/internal/service"
	"users/pkgs/logger"

	"github.com/go-chi/chi/v5"
)

type UserH struct {
	userservice service.UserS
	logger      logger.Logger
}

func NewUserH(us service.UserS, log logger.Logger) *UserH {
	return &UserH{
		userservice: us,
		logger:      log,
	}
}

func (h *UserH) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to create user")
	var req dto.UsercreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorf("invalid body request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		h.logger.Errorf("invalid body request")
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
	h.logger.Infof("user created successfully")
}

func (h *UserH) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to get users")
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
	h.logger.Infof("users fetched successfully")
}

func (h *UserH) GetUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to get user")
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
	h.logger.Infof("user fetched successfully")
}

func (h *UserH) UpdateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to update user")
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
	h.logger.Infof("user updated successfully")
}

func (h *UserH) DeleteUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to delete user")
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
	h.logger.Infof("user deleted successfully")
}

func (h *UserH) EnableAccount(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to enable account")
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
	h.logger.Infof("user enabled successfully")
}

func (h *UserH) LoginUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("received request to login user")
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userservice.LoginUser(r.Context(), &req)
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
		Message: "🔥 Welcome ",
		Data:    user,
	})
	h.logger.Infof("user logged in successfully")

}
