package main

import "testing"

func TestMysum(t *testing.T){

	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{21,21}, 32},
		test{[]int{3,4,5}, 32},
		test{[]int{1,1}, 32},
		test{[]int{-1,0,1}, 32},

	}
	for _,v := range tests {
		x := mysum(v.data...)
		if x != 3 {
			t.Error("Expected ",v.answer," Got ", x)
		}
	}
}

