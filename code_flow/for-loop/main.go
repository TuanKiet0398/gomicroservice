package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("-----------------infinite loop--------------")
	count := 0
	for {
		fmt.Println("counter", count)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println("-----------------skipping-------------------")

	for i:=1; i <=10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------------array-------------------")
	items := []string{"apple", "banana", "cherry"}
	for _, item := range items {
		fmt.Printf(" Item: %s\n", item)
	}
}