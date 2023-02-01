// Demonstrate errors
package main

import (
	"errors"
	"fmt"
	"log"
)

func area(length float64, width float64) (float64, error) {
	//check the if lenght is greater than 0
	if length <0 {
		err := errors.New("You cannot have negative lenght")
		return -1, err
	}
	if width < 0 {
		err:= errors.New("You cannot have any negative width")
	return-1, err
	}
result := length *width
return result, nil

}
func main() {
	lenght := 5.5
	width := 10
	//call the area function

	result, err := area(lenght, float64(width))
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}