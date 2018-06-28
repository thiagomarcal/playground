package main

import (
	"fmt"
)

func main() {
	result := stepPerms(35)
	fmt.Println(result)
}

func stepPerms(n int32) int32 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}

	a := int32(1)
	b := int32(1)
	c := int32(2)
	res := int32(0)

	for i := int32(3); i <= n; i++ {
		res = c + b + a
		a = b
		b = c
		c = res
	}

	return res
}
