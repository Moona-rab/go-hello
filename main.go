package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Enter width, length of wall, calculate area
// Add another wall y/n? while y (loop), else sum all walls
// enter efficiency of paint per litre
// area/efficiency = total litres of paint
// Functionise!

func length(scanner *bufio.Scanner) string {

	fmt.Print("\nPlease enter the length of your wall in metres: ")
	scanner.Scan()
	length_Str := scanner.Text()
	return length_Str
}

func height(scanner *bufio.Scanner) string {

	fmt.Print("\nPlease enter the height of your wall in metres: ")
	scanner.Scan()
	height_Str := scanner.Text()
	return height_Str
}

func walls(scanner *bufio.Scanner) string {

	fmt.Print("\nWould you like to add another wall (Y/N)? ")
	scanner.Scan()
	walls_Str := scanner.Text()
	return walls_Str

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	surf_area := []float64{}
	n := 0.0

	fmt.Println("\nWelcome to THE HAPPY PAINT CALCULATOR. We are here to help you with your paint project!")

	for {

		length_Str := length(scanner)

		user_length, err := strconv.ParseFloat(length_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", length_Str)
			return
		} else {
			fmt.Println(user_length, "m")
		}

		height_Str := height(scanner)

		user_height, err := strconv.ParseFloat(height_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", height_Str)
			return
		} else {
			fmt.Println(user_height, "m")
		}

		calc := (user_length * user_height)
		surf_area = append(surf_area, calc)

		walls_Str := walls(scanner)

		user_walls := strings.ToUpper(walls_Str)

		if user_walls == "Y" {

			continue

		} else if user_walls == "N" {

			for _, total := range surf_area {

				n += total
			}

			fmt.Printf("\nThe total surface area for paint is %.2f square-metres.", n)

		} else {
			fmt.Println("This is not a valid input")
			return
		}

		fmt.Println("\n\nEnter the efficacy of your paint in metres-squared per litre: ")
		scanner.Scan()
		efficacy_Str := scanner.Text()

		efficacy, err := strconv.ParseFloat(efficacy_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", efficacy_Str)
			return

		} else {

			paint_needed := n / efficacy
			fmt.Printf("\nYou will need %.2f litres of paint for %.2f square-metres of wall.", paint_needed, n)
			fmt.Println("\n\nThank you for using THE HAPPY PAINT CALCULATOR today. Happy painting!")
			break
		}

	}

}
