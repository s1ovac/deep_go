package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SingletonService struct {
	NotEmptyStruct bool
}

type UserService struct {
	NotEmptyStruct bool
}

type MessageService struct {
	NotEmptyStruct bool
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})
	container.RegisterSingletonType("SingletonService", func() any {
		return &SingletonService{}
	})

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)

	singletonService1, err := container.Resolve("SingletonService")
	assert.NoError(t, err)
	singletonService2, err := container.Resolve("SingletonService")
	assert.NoError(t, err)

	s1 := singletonService1.(*SingletonService)
	s2 := singletonService2.(*SingletonService)
	assert.True(t, s1 == s2)
}
