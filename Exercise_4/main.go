package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fn_1 := 0
	fn := 1
	return func() (old_fn_1 int) {
		old_fn_1, fn_1, fn = fn_1, fn, fn+fn_1
		return old_fn_1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
