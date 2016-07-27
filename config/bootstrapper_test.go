package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootstrap(t *testing.T) {
	b := Bootstrap()
	assert.IsType(t, &Config{}, b.Cfg)
}
