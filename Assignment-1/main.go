package main

import (
	"bufio"
	"fmt"
	"os"
)

type address struct{
	HouseNumber int
	Street string
	City string
	PinNumber int
}

type emp struct {

	EmpID int
	EmpName string
	EmpSalary float64
	EmpAddress address
}

type animal interface{
	Eat() string
	Sleep() string
	Breath() string
}

func (e emp) Eat() string{
	return "Employee is eating."
}
func (e emp) Sleep() string{
	return "Employee is sleeping."
}
func (e emp) Breath() string{
	return "Employee is breathing."
}

func main() {

	employee := (make([]animal,5,5))
	i:= 0

	for{
		fmt.Println("Enter I to add new employee or Q to quit or P to print employees details.")
		var choice string
		fmt.Scanln(&choice)

		if choice == "I"{
			fmt.Println("Please enter employee id:-")
			var empid int
			fmt.Scanln(&empid)
			fmt.Println("Please enter employee name:-")
			reader := bufio.NewReader(os.Stdin)
			var empname string
			empname, _ = reader.ReadString('\n')
			fmt.Println("Please enter employee salary:-")
			var empsal float64
			fmt.Scanln(&empsal)
			fmt.Println("Please enter employee house number:-")
			var emphouse int
			fmt.Scanln(&emphouse)
			fmt.Println("Please enter employee street:-")
			var empstreet string
			empstreet, _ = reader.ReadString('\n')
			fmt.Println("Please enter employee city:-")
			var empcity string
			empcity, _ = reader.ReadString('\n')
			fmt.Println("Please enter employee pincode:-")
			var emppin int
			fmt.Scanln(&emppin)

			empadd := address{HouseNumber:emphouse,Street: empstreet,City: empcity,PinNumber: emppin}

			empentry:= emp{EmpID: empid,EmpName: empname,EmpSalary: empsal,EmpAddress: empadd}

			if i < 5 {
				employee[i]=empentry
			}else{
				employee = append(employee,empentry)
			}
			fmt.Println("To show working of interface methods:-")
			fmt.Println(employee[i].Eat())
			fmt.Println(employee[i].Breath())
			fmt.Println(employee[i].Sleep())
			i = i+1
		} else if choice=="P"{
			fmt.Println(employee)

		}else{
			os.Exit(1)
		}
	}
}

