package exo1

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10+39+94",
			args: args{a: 10, b: 39, c: 94},
			want: 143,
		},
		{
			name: "-5+2+3",
			args: args{a: -5, b: 2, c: 3},
			want: 0,
		},
		{
			name: "19485928+39849834-49288346",
			args: args{a: 19485928, b: 39849834, c: -49288346},
			want: 10047416,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	type args struct {
		s    string
		char rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty",
			args: args{
				s:    "",
				char: 'a',
			},
			want: 0,
		},
		{
			name: "One occurrence",
			args: args{
				s:    "hello",
				char: 'h',
			},
			want: 1,
		},
		{
			name: "Two occurrences",
			args: args{
				s:    "hello",
				char: 'l',
			},
			want: 2,
		},
		{
			name: "No occurrences",
			args: args{
				s:    "hello",
				char: 'x',
			},
			want: 0,
		},
		{
			name: "Sentence",
			args: args{
				s:    "a man a plan a canal panama",
				char: 'a',
			},
			want: 10,
		},
		{
			name: "Long word",
			args: args{
				s:    "abracadabra",
				char: 'a',
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOccurrences(tt.args.s, tt.args.char); got != tt.want {
				t.Errorf("CountOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "10",
			args: args{n: 10},
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEven(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				numbers: []int{},
			},
			want: []int{},
		},
		{
			name: "Simple",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "List",
			args: args{
				numbers: []int{1, 23, 34, 119, 548, 4, 5, 8, 75, 1, 34, 65, 7, 3},
			},
			want: []int{34, 548, 4, 8, 34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterEven(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{a: 1},
			want: false,
		},
		{
			name: "10",
			args: args{a: 10},
			want: true,
		},
		{
			name: "66336",
			args: args{a: 66336},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.a); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxOfFour(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "One",
			args: args{a: 5, b: -19, c: 37, d: 25},
			want: 37,
		},
		{
			name: "Two",
			args: args{a: 30948, b: 409822, c: 304, d: 999999},
			want: 999999,
		},
		{
			name: "Negative",
			args: args{a: -30948, b: -409822, c: -304, d: -999999},
			want: -304,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxOfFour(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("MaxOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "One letter",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Hello",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "Hello 世界",
			args: args{
				s: "Hello 世界",
			},
			want: "界世 olleH",
		},
		{
			name: "Happy birthday",
			args: args{
				s: "Happy birthday ! 🎉🥳🎂",
			},
			want: "🎂🥳🎉 ! yadhtrib yppaH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
