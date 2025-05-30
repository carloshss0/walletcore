package gateway

import "github.com/carloshss0/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}