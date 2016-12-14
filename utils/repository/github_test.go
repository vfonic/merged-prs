package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	config := &ServiceConfig{
		Token: "foo",
		Owner: "bar",
		Repo:  "baz",
	}

	c := New(config)
	err := c.Connect()
	assert.NoError(t, err, "Communication with GitHub should be okay")
}
