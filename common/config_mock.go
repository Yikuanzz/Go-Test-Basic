package common

import (
	"github.com/stretchr/testify/mock"
)

type MockConfigService struct {
	mock.Mock
}

func (m *MockConfigService) GetConfig() Config {
	args := m.Called()
	return args.Get(0).(Config)
}
