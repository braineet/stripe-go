package ephemeralkey

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/braineet/stripe-go/v71"
	_ "github.com/braineet/stripe-go/v71/testing"
)

func TestEphemeralKeyDel(t *testing.T) {
	key, err := Del("ephkey_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, key)
}

func TestEphemeralKeyNew(t *testing.T) {
	key, err := New(&stripe.EphemeralKeyParams{
		Customer:      stripe.String("cus_123"),
		StripeVersion: stripe.String(stripe.APIVersion),
	})
	assert.Nil(t, err)
	assert.NotNil(t, key)
}
