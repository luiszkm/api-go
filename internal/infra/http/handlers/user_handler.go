package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/luiszkm/api/internal/Domain/entity"
	"github.com/luiszkm/api/internal/dto"
	"github.com/luiszkm/api/internal/infra/database"
)

type userHandler struct {
	UserDb        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewUserHandler(userDb database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiriesIn int) *userHandler {
	return &userHandler{
		UserDb:        userDb,
		Jwt:           jwt,
		JwtExpiriesIn: JwtExpiriesIn,
	}
}
func (h *userHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDb.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, err := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiriesIn)).Unix(),
	})
	if err != nil {
		fmt.Println("Erro na Codificação do JWT:", err)
		http.Error(w, "Erro Interno do Servidor", http.StatusInternalServerError)
		return
	}
	fmt.Println("Token String:", h.Jwt)

	accesToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesToken)
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDb.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
