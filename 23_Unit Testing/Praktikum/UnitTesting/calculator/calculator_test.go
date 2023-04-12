package calculator

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSubstrcation(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substraction(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("substract() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMultipy(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplication(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Division(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want = %v", got, tt.want)
			}
		})
	}
}
