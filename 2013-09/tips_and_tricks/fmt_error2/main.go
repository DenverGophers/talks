package main

// START OMIT
import (
	"fmt"
	"log"
)

func main() {
	if err := iDunBlowedUp(-100); err != nil {
		err = fmt.Errorf("Something went wrong: %s\n", err)
		log.Println(err)
		return
	}
	fmt.Printf("Success!")
}

func iDunBlowedUp(val int) error {
	return fmt.Errorf("invalid value %d", val)
}

// END OMIT
