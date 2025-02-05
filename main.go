package main

import (
	"fmt"; "github.com/SerhiiKhyzhko/puppy"
)

func main() {
	fmt.Println("hello there")
	firstDog := puppy.Bark()
	secondDog :=puppy.Barks()
	lst := [2]string{firstDog, secondDog}
	for _, dog:= range lst {
		fmt.Println(dog)
	}
}