package main

import "fmt"

func main() {
	fmt.Println("loops")

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// fmt.Println(days)

	// for d:=0;d<len(days);d++ {
	// 	fmt.Println(days[d])
	// }

	// for _, day := range days {
	// 	fmt.Println(day)
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	orphanVal := 1

	for orphanVal < 10 {

		if orphanVal == 3 {
			goto hruday
		}

		if orphanVal == 6 {
			orphanVal++
			// continue
			break
		}

		fmt.Println("val is : ", orphanVal)
		orphanVal++
	}

	hruday:
		fmt.Println("hruday is here")
}
