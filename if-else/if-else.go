package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 || i%6 == 0 || i%9 == 0 { //Checking if number divded by 3,6,9...
			fmt.Println(i)
		} else {
			continue //Skip if not fall short of conditions
		}
	}

}
