package main

import "fmt"

type dog struct {
	Breed string
	Weight int

}


func main(){

	dog1:=dog{Breed: "beagle",Weight: 25}
	fmt.Println(dog1)
	fmt.Println(dog1.Weight)
}
