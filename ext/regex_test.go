package ext_test

import (
	"github.com/gopkg-dev/go-common/ext"
	"reflect"
	"testing"
)

func TestMatchOneOf(t *testing.T) {
	type args struct {
		patterns []string
		text     string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal test",
			args: args{
				patterns: []string{`aaa(\d+)`, `hello(\d+)`},
				text:     "hello12345",
			},
			want: []string{
				"hello12345", "12345",
			},
		},
		{
			name: "normal test",
			args: args{
				patterns: []string{`aaa(\d+)`, `bbb(\d+)`},
				text:     "hello12345",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ext.MatchOneOf(tt.args.text, tt.args.patterns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchOneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchAll(t *testing.T) {
	type args struct {
		pattern string
		text    string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "normal test",
			args: args{
				pattern: `hello(\d+)`,
				text:    "hello12345hello123",
			},
			want: [][]string{
				{
					"hello12345", "12345",
				},
				{
					"hello123", "123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ext.MatchAll(tt.args.text, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
