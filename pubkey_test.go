package routinghelpers

import (
	"context"
	"testing"

	"github.com/chiangmaioneluv/go-libp2p/core/routing"
	"github.com/chiangmaioneluv/go-libp2p/core/test"
)

func TestGetPublicKey(t *testing.T) {
	t.Parallel()

	d := Parallel{
		Routers: []routing.Routing{
			Parallel{
				Routers: []routing.Routing{
					&Compose{
						ValueStore: &LimitedValueStore{
							ValueStore: new(dummyValueStore),
							Namespaces: []string{"other"},
						},
					},
				},
			},
			Tiered{
				Routers: []routing.Routing{
					&Compose{
						ValueStore: &LimitedValueStore{
							ValueStore: new(dummyValueStore),
							Namespaces: []string{"pk"},
						},
					},
				},
			},
			&Compose{
				ValueStore: &LimitedValueStore{
					ValueStore: new(dummyValueStore),
					Namespaces: []string{"other", "pk"},
				},
			},
			&struct{ Compose }{Compose{ValueStore: &LimitedValueStore{ValueStore: Null{}}}},
			&struct{ Compose }{},
		},
	}

	pid, _ := test.RandPeerID()

	ctx := context.Background()
	if _, err := d.GetPublicKey(ctx, pid); err != routing.ErrNotFound {
		t.Fatal(err)
	}
}
