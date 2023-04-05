package expirables_test

import (
	"github.com/bliink-io/expirables"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewExpirable(t *testing.T) {
	e := expirables.NewExpirable(func() interface{} {
		return "value"
	}, time.Second*5)

	assert.Equal(t, "value", e.Get())
}

func TestExpirableDuration(t *testing.T) {
	e := expirables.NewExpirable(func() interface{} {
		return "value"
	}, time.Second*5)

	assert.Equal(t, "value", e.Get())
	time.Sleep(time.Second * 6)
	assert.Equal(t, "value", e.Get())
}
