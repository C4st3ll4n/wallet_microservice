package client

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/C4st3ll4n/wallet/internal/gateway"
	"time"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUsecase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(gateway gateway.ClientGateway) *CreateClientUsecase {
	return &CreateClientUsecase{gateway}
}

func (uc *CreateClientUsecase) Execute(dto CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(dto.Name, dto.Email)

	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		client.ID,
		client.Name,
		client.Email,
		client.CreatedAt,
		client.UpdatedAt,
	}, nil
}
