package main

import "fmt"

func main() {
	birdCount := []int{4, 7, 3, 2, 1, 1, 2, 0, 2, 3, 2, 7, 1, 3, 0, 6, 5, 3, 7, 2, 3}
	totalWeek := 0
	for i := 0; i < len(birdCount); i += 7 {
		if ((i / 7) + 1) == 2 {
			for j := i; j < i+7; j++ {
				totalWeek += birdCount[j]
			}
		}
	}
	fmt.Println(totalWeek)
}
