package main

import "fmt"

func main() {
	workingWeek := map[int]string{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thi", 5: "Fri"}
	for i := 0; i < 10; i++ {
		m := ""
		for key, value := range workingWeek {
			m += fmt.Sprintf(`%d:"%s", `, key, value)
		}
		fmt.Println(m)
	}
	value, ok := workingWeek[42]
	fmt.Println(value, ok)
}
