package createtransaction

import (
	"context"
	"testing"

	"github.com/carloshss0/walletcore/internal/entity"
	"github.com/carloshss0/walletcore/internal/event"
	"github.com/carloshss0/walletcore/internal/usecase/mocks"
	"github.com/carloshss0/walletcore/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestCreateTransactionUseCaseExecute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "c1@email.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "c2@email.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo: account2.ID,
		Amount: 100,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransactionCreated := event.NewTransactionCreated()
	eventBalanceUpdated := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransactionCreated, eventBalanceUpdated)

	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNotCalled(t, "Do", 1)



	

}