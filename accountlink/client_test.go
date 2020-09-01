package accountlink

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/braineet/stripe-go/v71"
	_ "github.com/braineet/stripe-go/v71/testing"
)

func TestAccountLinkNew(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_123"),
		Collect:    stripe.String(string(stripe.AccountLinkCollectCurrentlyDue)),
		RefreshURL: stripe.String("https://stripe.com/refresh"),
		ReturnURL:  stripe.String("https://stripe.com/return"),
		Type:       stripe.String(string(stripe.AccountLinkTypeAccountOnboarding)),
	}
	link, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
