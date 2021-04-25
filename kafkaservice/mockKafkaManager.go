package kafkaservice

import (
	"context"

	"github.com/stretchr/testify/mock"
)

//MockKafkaSvc is the mock implementation of KafkaSvc
type MockKafkaSvc struct {
	mock.Mock
}

//WriteToKafka is the mock implementation
func (m *MockKafkaSvc) WriteToKafka(ctx context.Context, message interface{}) error {
	args := m.Mock.Called(ctx, message)
	return args.Error(0)
}
