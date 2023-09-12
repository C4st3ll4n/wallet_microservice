package gateway

import "github.com/C4st3ll4n/wallet/internal/entity"

type AccountGateway interface {
	Save(account entity.Account) error
	FindById(id string) (*entity.Account, error)
}
