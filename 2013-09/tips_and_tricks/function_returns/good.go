package main

import (
	"fmt"
	"log"
)

// START OMIT
func someFunction(val int) (bool, error) {
	if val == 0 {
		return false, nil
	}
	if val < 0 {
		return false, fmt.Errorf("value can't be negative %d", val)
	}
	return true, nil
}

// END OMIT
