package shortener

import (
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

