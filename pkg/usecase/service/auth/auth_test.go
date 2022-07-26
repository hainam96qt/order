package auth

import (
	"context"
	"order-gokomodo/pkg/entities"
	test_config "order-gokomodo/testing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthenticationService(t *testing.T) {
	ctx := context.Background()
	configTest, err := test_config.NewConfigTest()
	assert.NoError(t, err)

	service, err := NewAuthenticationService(configTest)
	assert.NoError(t, err)

	err = service.Query.TruncateData(ctx)
	assert.NoError(t, err)

	result, err := service.Register(ctx, &entities.RegisterRequest{
		Email:    "email01",
		Password: "password01",
		Name:     "name01",
		Role:     "seller",
		Address:  "This is an address",
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)

	result, err = service.Register(ctx, &entities.RegisterRequest{
		Email:    "email02",
		Password: "password02",
		Name:     "name01",
		Role:     "buyer",
		Address:  "This is an address",
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)

	result, err = service.Register(ctx, &entities.RegisterRequest{
		Email:    "email02",
		Password: "password02",
		Name:     "name01",
		Role:     "anonymous",
		Address:  "This is an address",
	})
	assert.Error(t, err)
	assert.Nil(t, result)
}
