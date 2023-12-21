package urlcoder

import (
	"testing"
)

func TestNewDecodeAlias(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		id    int
	}{
		{
			name:  "alias = cb, id = 21",
			alias: "cd",
			id:    21,
		},
		{
			name:  "alias = e9a, id = 19158",
			alias: "e9a",
			id:    19158,
		},
		{
			name:  "alias = cb, id = 21",
			alias: "b",
			id:    1,
		},
		{
			name:  "alias = cb, id = 21",
			alias: "google.com",
			id:    57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Add tests
		})
	}
}
