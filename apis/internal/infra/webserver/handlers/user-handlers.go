package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/dto"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/schemas"
	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/util"
)

type UserHandler struct {
	UserService database.UserInterface
}

func NewUserHandler(service database.UserInterface) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
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
