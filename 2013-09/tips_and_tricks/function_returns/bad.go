package main

import (
	"fmt"
	"log"
)

// START OMIT
func someFunction(val int) (ok bool, err error) {
	if val == 0 {
		return false, nil
	}
	if val < 0 {
		return false, fmt.Errorf("value can't be negative %d", val)
	}
	ok = true
	return
}

// END OMIT
