package web

import (
	"encoding/json"
	"github.com/C4st3ll4n/wallet/internal/usecase/account"
	"net/http"
)

type WebAccountHandler struct {
	CreateAccountUsecase account.CreateAccountUsecase
}

func NewWebAccountHandler(usecase account.CreateAccountUsecase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUsecase: usecase,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto account.CreateAccountInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateAccountUsecase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
