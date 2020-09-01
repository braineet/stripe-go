package mandate

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	_ "github.com/braineet/stripe-go/v71/testing"
)

func TestMandateMethodGet(t *testing.T) {
	pm, err := Get("mandate_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}
