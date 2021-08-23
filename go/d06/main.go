package main

import "fmt"

func out(s string) {
	var s1, s2 string
	for ind, c := range s {
		// fmt.Println(ind)
		if ind%2 == 0 {
			s1 += string(c)
			// fmt.Println("s1: "+ s1)
		} else {
			s2 += string(c)
			// fmt.Println(s2)
		}

	}
	fmt.Println(s1 + " " + s2)
}
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	var S string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&S)
		out(S)
	}
}
