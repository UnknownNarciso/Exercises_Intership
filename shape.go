package main

import "fmt"

// 13/05/2022
func main() {
	var height = 0
	fmt.Println("Hey user this program is used to create diamond pattern\n   ")
	fmt.Println("You need to select height you want for the pattern need to be 3 to 45 because its group of numbers can do perfect pattern of diamond\n  ")
	fmt.Println("Input height of pattern of diamond : ")
	fmt.Scanln(&height)
	for i := 1; i <= height; i += 2 {
		for j := 1; j < height-i/2; j++ {

			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := height - 2; i > 0; i -= 2 {
		for j := 1; j < height-i/2; j++ {

			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println()
	}

}
