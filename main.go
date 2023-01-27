package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Enter width, length of wall, calculate area
// Add another wall y/n?  if y, loop, else sum all walls
// enter no. doors
// enter no. of coats
// enter efficiency of paint per litre
// (area/efficiency - doors area) * coats = total litres of paint

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

func doors(scanner *bufio.Scanner) string {

	fmt.Print("\nPlease enter the number of doors: ")
	scanner.Scan()
	doors_Str := scanner.Text()
	return doors_Str
}

func coats(scanner *bufio.Scanner) string {

	fmt.Println("\nHow many coats of paint will you apply? ")
	scanner.Scan()
	coats_Str := scanner.Text()
	return coats_Str
}

func efficacy(scanner *bufio.Scanner) string {

	fmt.Println("\nEnter the efficacy of your paint in metres-squared per litre: ")
	scanner.Scan()
	efficacy_Str := scanner.Text()
	return efficacy_Str
}

func smiley() {

	fmt.Println(`
                     ***                  ***
                    *****                *****
                    *****                *****
                     ***                  ***
          ***                                        ***
           ***                                      ***
            ***                                    ***
             ***                                  ***
               ***                              ***
                 ***                          ***
                   ***                      ***
                      **********************
                         ****************
	
	`)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	surf_area := []float64{}
	n := 0.0

	fmt.Println("\nWelcome to THE HAPPY PAINT CALCULATOR. We are here to help you with your paint project!")
	smiley()

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

			fmt.Printf("\nThe total surface area for paint is %.2f square-metres.\n", n)

		} else {
			fmt.Println("This is not a valid input")
			return
		}

		doors_Str := doors(scanner)

		user_doors, err := strconv.ParseFloat(doors_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", doors_Str)
			return
		} else {
			fmt.Println(user_doors, "door(s)")
		}

		doors_area := user_doors * 2

		coats_Str := coats(scanner)

		user_coats, err := strconv.ParseFloat(coats_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", coats_Str)
			return
		} else {
			fmt.Println(user_coats, "coats")
		}

		efficacy_Str := efficacy(scanner)

		efficacy, err := strconv.ParseFloat(efficacy_Str, 64)

		if err != nil {
			fmt.Printf("'%s' is not a number.", efficacy_Str)
			return

		} else {

			paint_needed := (n/efficacy - doors_area) * user_coats
			fmt.Printf("\nYou will need %.2f litres of paint for %.2f square-metres of wall.", paint_needed, n)
			fmt.Println("\n\nThank you for using THE HAPPY PAINT CALCULATOR today. Happy painting!")
			smiley()
			break
		}

	}

}
