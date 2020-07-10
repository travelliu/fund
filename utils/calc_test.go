// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package utils

import (
	"testing"
)

func TestCalcFloat64(t *testing.T) {
	type args struct {
		f   float64
		pos int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "107179.01843999999",
			args: args{
				f:   107179.01843999999,
				pos: 2,
			},
			want: 107179.02,
		},
		{
			name: "107179.01443999999",
			args: args{
				f:   107179.01443999999,
				pos: 2,
			},
			want: 107179.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcFloat64(tt.args.f, tt.args.pos); got != tt.want {
				t.Errorf("CalcFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
