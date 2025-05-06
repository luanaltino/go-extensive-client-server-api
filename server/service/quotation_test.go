package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindExchangeWithSuccess(t *testing.T) {
	result, err := FindExchange()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.USDBRL.Code)
}
