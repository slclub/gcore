package gcore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnginFake(t *testing.T) {
	en := New()

	assert.Equal(t, "127.0.0.1:8080", en.http_addr.String())

	en.DriverRegister()
}
