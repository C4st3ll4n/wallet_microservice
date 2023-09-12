package gateway

import "github.com/C4st3ll4n/wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction entity.Transaction) error
}
