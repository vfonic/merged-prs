package config

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefault(t *testing.T) {
	usr, _ := user.Current()
	c := NewDefault()

	assert.Equal(t, usr.Username, c.User)
	assert.Equal(t, usr.HomeDir, c.Home)
}

func TestMock(t *testing.T) {
	c := NewMock()
	assert.Equal(t, "test", c.User)
	assert.Equal(t, ".", c.Home)
}
