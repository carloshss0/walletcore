package createaccount

import (
	"testing"

	"github.com/carloshss0/walletcore/internal/entity"
	"github.com/carloshss0/walletcore/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestCreateAccountUseCaseExecute(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.Mock.On("Get", client.ID).Return(client, nil)
	

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)

	output, err := uc.Execute(CreateAccountInputDTO{
		ClientID: client.ID,
	})

	assert.NotNil(t, output)
	assert.Nil(t, err)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}


