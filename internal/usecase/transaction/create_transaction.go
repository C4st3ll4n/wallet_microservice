package transaction

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/C4st3ll4n/wallet/internal/gateway"
	"github.com/C4st3ll4n/wallet/pkg/events"
)

type CreateTransactionUsecase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
	eventDispatcher    events.EventDispatcherInterfcae
	transaction        events.EventInterface
}

func NewCreateTransaction(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	EventDispatcher events.EventDispatcherInterfcae,
	Transaction events.EventInterface,
) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		eventDispatcher:    EventDispatcher,
		transaction:        Transaction,
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

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	uc.transaction.SetPayload(output)
	err = uc.eventDispatcher.Dispatch(uc.transaction)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateTransactionInputDTO struct {
	IDAccountFrom        string
	IDAccountDestination string
	Amount               float64
}

type CreateTransactionOutputDTO struct {
	ID string
}
