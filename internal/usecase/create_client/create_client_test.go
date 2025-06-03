package createclient

import (
	"testing"

	"github.com/carloshss0/walletcore/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestCreateClientUseCaseExecute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name: "John",
		Email: "john@email.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)
	assert.Equal(t, "John", output.Name)
	assert.Equal(t, "john@email.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

