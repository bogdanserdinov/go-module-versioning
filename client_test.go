package sdk_versioning

import (
	"context"
	"testing"
)

func TestCoingeckoClient(t *testing.T) {
	client := NewClient()

	ctx := context.Background()

	client.(ctx, "bitcoin")
}
