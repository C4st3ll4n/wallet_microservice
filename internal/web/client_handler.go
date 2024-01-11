package web

import (
	"encoding/json"
	"github.com/C4st3ll4n/wallet/internal/usecase/client"
	"net/http"
)

type WebClientHandler struct {
	CreateClientUsecase client.CreateClientUsecase
}

func NewWebClientHandler(usecase client.CreateClientUsecase) *WebClientHandler {
	return &WebClientHandler{CreateClientUsecase: usecase}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto client.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateClientUsecase.Execute(dto)
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
