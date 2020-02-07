package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	type args struct {
		i int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1の時1を返す", args: args{i: 1}, want: "1"},
		{name: "3の倍数の時Fizzを返す", args: args{i: 297}, want: "Fizz"},
		{name: "5の倍数の時Buzzを返す", args: args{i: 535}, want: "Buzz"},
		{name: "3の倍数かつ5の倍数の時FizzBuzzを返す", args: args{i: 300}, want: "FizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.i); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
