package main

import (
	"fmt"
	"log"
)

// START OMIT
func someFunction(val int) (ok bool, err error) {
	if val == 0 {
		return
	}
	if val < 0 {
		err = fmt.Errorf("value can't be negative %d", val)
		return
	}
	ok = true
	return
}

// END OMIT
