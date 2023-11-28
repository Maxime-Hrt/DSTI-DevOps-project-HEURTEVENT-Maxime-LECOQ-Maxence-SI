package mocks

import (
	"devops-project/src"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockRedisService struct {
	mock.Mock
}

func (m *MockRedisService) Set(key string, value interface{}, expiration time.Duration) error {
	args := m.Called(key, value, expiration)
	return args.Error(0)
}

func (m *MockRedisService) SaveContactInRedis(contact *src.Contact) error {
	args := m.Called(contact)
	return args.Error(0)
}

func (m *MockRedisService) GetContactFromRedis(id string) (*src.Contact, error) {
	args := m.Called(id)
	return args.Get(0).(*src.Contact), args.Error(1)
}

func (m *MockRedisService) GetAllContactsFromRedis() ([]src.Contact, error) {
	args := m.Called()
	return args.Get(0).([]src.Contact), args.Error(1)
}

func (m *MockRedisService) DeleteContactFromRedis(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRedisService) UpdateContactInRedis(id string, updatedContact *src.Contact) error {
	args := m.Called(id, updatedContact)
	return args.Error(0)
}
