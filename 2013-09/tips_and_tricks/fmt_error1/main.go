package main

// START OMIT
import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := iDunBlowedUp(-100); err != nil {
		err = errors.New(fmt.Sprintf("Something went wrong: %s\n", err))
		log.Println(err)
		return
	}
	fmt.Printf("Success!")
}

func iDunBlowedUp(val int) error {
	return errors.New(fmt.Sprintf("invalid value %d", val))
}

// END OMIT
