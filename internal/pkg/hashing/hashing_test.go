package hashing_test

import (
	"testing"

	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/hashing"
	"github.com/stretchr/testify/assert"
)

func TestHashStringKey(t *testing.T) {
	type args struct {
		key     string
		buckets int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "invalid buckets: 0",
			args: args{key: "foo", buckets: 0},
			want: -1,
		},
		{
			name: "invalid buckets: -1",
			args: args{key: "foo", buckets: -1},
			want: -1,
		},
		{
			name: "valid key: foo",
			args: args{key: "foo", buckets: 3},
			want: 1,
		},
		{
			name: "valid key: bar",
			args: args{key: "bar", buckets: 3},
			want: 0,
		},
		{
			name: "valid key: baz",
			args: args{key: "baz", buckets: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hashing.HashStringKey(tt.args.key, tt.args.buckets))
		})
	}
}
