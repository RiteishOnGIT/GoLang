package AcDc

func Sum(xi ...int) int {
	x := 0
	for _,v := range xi{
		x = x + v
	}
	return x
}

