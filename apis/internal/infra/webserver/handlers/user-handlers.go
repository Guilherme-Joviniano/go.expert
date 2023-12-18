package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/dto"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/schemas"
	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/util"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserService  database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(service database.UserInterface, jwtConfig *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserService:  service,
		Jwt:          jwtConfig,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetToken(
	w http.ResponseWriter,
	req *http.Request,
) {
	authenticationInput, err := util.RequestToTypeAdapter[dto.AuthenticationInput](req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserService.FindByEmail(authenticationInput.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(authenticationInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, err := h.Jwt.Encode(map[string]interface{}{
		"sub": user.Id.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (h *UserHandler) CreateUser(
	w http.ResponseWriter,
	req *http.Request,
) {
	createUserInput, err := util.RequestToTypeAdapter[dto.CreateUserInput](req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(createUserInput.Name, createUserInput.Email, createUserInput.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserService.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
