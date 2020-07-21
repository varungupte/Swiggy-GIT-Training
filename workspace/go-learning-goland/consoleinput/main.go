package main

import (
	"fmt"
	"sort"
)

func main(){
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("---------------------")
	//
	//fmt.Println("Enter the salary: ")
	//
	//salary1, _ := reader.ReadString('\n')
	//fmt.Println("Salary ",salary1)
	//salary, _ := strconv.Atoi(strings.TrimSpace(salary1))
	//fmt.Printf("Salary : %d \n", salary)
	//
	//fmt.Print("Enter number of month: ")
	//month1, _ := reader.ReadString('\n')
	//month, _ := strconv.Atoi(strings.TrimSpace(month1))

	states:= make(map[string] string)

	states["1"]="one"
	states["2"]="two"

	keys := make([]string, 0, len(states))
	values := make([]string, 0, len(states))

	for k, v := range states {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Println(keys)


	ints:= make([]int,5,5)

	ints[0]=23
	ints[1]=98
	ints[2]=2
	ints[3]=45
	ints[4]=0

	fmt.Println(ints)

	//sort.Ints(ints)
	//sort.Reverse()

	//sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	ints2 := ints[1:3]
	fmt.Println(ints2)

	sort.Sort(sort.IntSlice(ints2))
	fmt.Println(ints2)

	//sort.Slice(ints, func(i, j int) bool {
	//	return ints[i] > ints[j]
	//})

	//fmt.Println(ints)


	//sort.Ints(ints)
	//fmt.Println(ints)
	//
	//
	//total := salary*month
	//
	//fmt.Printf("Total is %d ",total)



}
