package ext_test

import (
	"github.com/gopkg-dev/go-common/ext"
	"testing"
)

func TestLimitLength(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				s:      "你好 hello",
				length: 8,
			},
			want: "你好 hello",
		},
		{
			name: "normal test",
			args: args{
				s:      "你好 hello",
				length: 6,
			},
			want: "你好 ...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ext.LimitLength(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("LimitLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				name: "hello/world",
			},
			want: "hello world",
		},
		{
			name: "normal test",
			args: args{
				name: "hello:world",
			},
			want: "hello：world",
		},
		{
			name: "overly long strings test",
			args: args{
				name: "super 超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长", // length 81
			},
			want: "super 超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级长超级...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ext.FileName(tt.args.name, ""); got != tt.want {
				t.Errorf("FileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
