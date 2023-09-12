package account

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/C4st3ll4n/wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ClientID string
}

type CreateAccountUsecase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUsecase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUsecase {
	return &CreateAccountUsecase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (usecase *CreateAccountUsecase) Execute(dto CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := usecase.ClientGateway.Get(dto.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = usecase.AccountGateway.Save(*account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ClientID: account.ID,
	}, nil
}
