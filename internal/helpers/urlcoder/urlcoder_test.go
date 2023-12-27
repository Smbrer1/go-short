package urlcoder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Smbrer1/go-short/internal/helpers/urlcoder"
)

func TestNewEncodeAlias(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		id    int64
	}{
		{
			name:  "alias=cb;id=21",
			alias: "L",
			id:    21,
		},
		{
			name:  "alias=e9a;id=19158",
			alias: "4z0",
			id:    19158,
		},
		{
			name:  "alias=cb;id=21",
			alias: "a",
			id:    36,
		},
		{
			name:  "alias=cb;id=21",
			alias: "v",
			id:    57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alias, _ := urlcoder.Encode(tt.id)
			assert.Equal(t, tt.alias, alias)
		})
	}
}



func TestNewDecodeAlias(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		id    int64
	}{
		{
			name:  "alias=cb;id=21",
			alias: "L",
			id:    21,
		},
		{
			name:  "alias=e9a;id=19158",
			alias: "4z0",
			id:    19158,
		},
		{
			name:  "alias=cb;id=21",
			alias: "a",
			id:    36,
		},
		{
			name:  "alias=cb;id=21",
			alias: "v",
			id:    57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, _ := urlcoder.Decode(tt.alias)
			assert.Equal(t, tt.id, id)
		})
	}
}
