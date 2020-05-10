package ext_test

import (
	"github.com/gopkg-dev/go-common/ext"
	"testing"


)

func TestDomain(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				url: "http://www.aa.com",
			},
			want: "aa",
		},
		{
			name: "normal test",
			args: args{
				url: "https://aa.com",
			},
			want: "aa",
		},
		{
			name: "normal test",
			args: args{
				url: "aa.cn",
			},
			want: "aa",
		},
		{
			name: "normal test",
			args: args{
				url: "www.aa.cn",
			},
			want: "aa",
		},
		{
			name: ".com.cn test",
			args: args{
				url: "http://www.aa.com.cn",
			},
			want: "aa",
		},
		{
			name: "Universal test",
			args: args{
				url: "http://aa",
			},
			want: "Universal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ext.Domain(tt.args.url); got != tt.want {
				t.Errorf("Domain() = %v, want %v", got, tt.want)
			}
		})
	}
}
