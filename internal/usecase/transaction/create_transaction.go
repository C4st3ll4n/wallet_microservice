package transaction

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/C4st3ll4n/wallet/internal/gateway"
)

type CreateTransactionUsecase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransaction(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUsecase) Execute(dto CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindById(dto.IDAccountFrom)

	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindById(dto.IDAccountDestination)

	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, dto.Amount)

	if err != nil {
		return nil, err
	}

	//transaction.Commit()

	err = uc.TransactionGateway.Create(*transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}

type CreateTransactionInputDTO struct {
	IDAccountFrom        string
	IDAccountDestination string
	Amount               float64
}

type CreateTransactionOutputDTO struct {
	ID string
}
