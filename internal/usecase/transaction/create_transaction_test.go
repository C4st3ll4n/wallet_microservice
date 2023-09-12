package transaction

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "email@mail.com")
	client2, _ := entity.NewClient("Client 2", "mail@email.com")

	account1 := entity.NewAccount(client1)
	account2 := entity.NewAccount(client2)

	account1.Credit(1500)
	account2.Credit(500)

	mockTransaction := TransactionGatewayMock{}
	mockAccount := AccountGatewayMock{}

	mockAccount.On("FindById", account1.ID).Return(account1, nil)
	mockAccount.On("FindById", account2.ID).Return(account2, nil)
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDTO := CreateTransactionInputDTO{
		IDAccountFrom:        account2.ID,
		IDAccountDestination: account1.ID,
		Amount:               500,
	}

	usecase := NewCreateTransaction(&mockTransaction, &mockAccount)
	output, err := usecase.Execute(inputDTO)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)

	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
}
