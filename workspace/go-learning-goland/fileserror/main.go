package main

import(
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("Sample.txt") // For read access.
	if err == nil{
		fmt.Println(file)
	}else{
		fmt.Println(err.Error())
	}
}