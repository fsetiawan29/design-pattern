package adapter

import (
	"reflect"
	"testing"
)

func TestNewRoundHole(t *testing.T) {
	type args struct {
		radius float32
	}
	tests := []struct {
		name string
		args args
		want *RoundHole
	}{
		{
			name: "Default Test",
			args: args{
				radius: 13,
			},
			want: &RoundHole{
				radius: 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoundHole(tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoundHole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundHole_getRadius(t *testing.T) {
	type fields struct {
		radius float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "Default Test",
			fields: fields{
				radius: 13,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &RoundHole{
				radius: tt.fields.radius,
			}
			if got := rh.getRadius(); got != tt.want {
				t.Errorf("RoundHole.getRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundPeg_getRadius(t *testing.T) {
	type fields struct {
		radius float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "Default Test",
			fields: fields{
				radius: 13,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &RoundPeg{
				radius: tt.fields.radius,
			}
			if got := rp.getRadius(); got != tt.want {
				t.Errorf("RoundPeg.getRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquarePegAdapter_getRadius(t *testing.T) {
	type fields struct {
		SquarePeg *SquarePeg
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "Default Test",
			fields: fields{
				SquarePeg: &SquarePeg{
					width: 14,
				},
			},
			want: 9.899495,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spa := &SquarePegAdapter{
				SquarePeg: tt.fields.SquarePeg,
			}
			if got := spa.getRadius(); got != tt.want {
				t.Errorf("SquarePegAdapter.getRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundHole_Fits(t *testing.T) {
	type fields struct {
		radius float32
	}
	type args struct {
		roundPeg RoundPegItf
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Default Test",
			fields: fields{
				radius: 5,
			},
			args: args{
				roundPeg: &RoundPeg{
					radius: 10,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := RoundHole{
				radius: tt.fields.radius,
			}
			if got := rh.Fits(tt.args.roundPeg); got != tt.want {
				t.Errorf("RoundHole.Fits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRoundPeg(t *testing.T) {
	type args struct {
		radius float32
	}
	tests := []struct {
		name string
		args args
		want *RoundPeg
	}{
		{
			name: "Default Test",
			args: args{
				radius: 5,
			},
			want: &RoundPeg{
				radius: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoundPeg(tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoundPeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSquarePeg(t *testing.T) {
	type args struct {
		width float32
	}
	tests := []struct {
		name string
		args args
		want *SquarePeg
	}{
		{
			name: "Default Test",
			args: args{
				width: 10,
			},
			want: &SquarePeg{
				width: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSquarePeg(tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSquarePeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSquarePegAdapter(t *testing.T) {
	type args struct {
		squarePeg *SquarePeg
	}
	tests := []struct {
		name string
		args args
		want *SquarePegAdapter
	}{
		{
			name: "Default Test",
			args: args{
				squarePeg: &SquarePeg{
					width: 5,
				},
			},
			want: &SquarePegAdapter{
				SquarePeg: &SquarePeg{
					width: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSquarePegAdapter(tt.args.squarePeg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSquarePegAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}
