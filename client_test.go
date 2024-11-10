package sdk_versioning_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bogdanserdinov/sdk-versioning-test"
)

func TestCoingeckoClient(t *testing.T) {
	client := sdk_versioning.NewClient()

	ctx := context.Background()

	price, err := client.GetPrice(ctx, "bitcoin")
	require.NoError(t, err, "unable to get price")

	assert.Greater(t, price.Bitcoin.Usd, 0, "price should be greater than 0")
}
