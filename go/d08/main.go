package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	var s1, s2 string
	var I string
	var PB = make(map[string]string)
	var q = make(map[int]string)
	// var r = make(map[int]string)

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scanln(&s1, &s2)
		PB[s1] = s2
	}

	for i := 0; i < 100000; i++ {
		_, err := fmt.Scan(&I)
		if err != nil {
			break
		}
		q[i] = I
	}
	for i := 0; i < len(q); i++ {
		if v, found := PB[q[i]]; found {
			fmt.Printf("%s=%s\n", q[i], v)
		} else {
			fmt.Println("Not found")
		}
	}

}
