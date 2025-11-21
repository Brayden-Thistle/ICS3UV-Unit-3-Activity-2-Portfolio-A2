/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-19
* @fileoverfiew this program asks you for the Provinces and Territories of Canada and then displays them
*/

package main

import (
"fmt"
"strings"
"bufio"
"os"
)

func main () {
	//variables
	var province1 string
	var province2 string
	var province3 string
	var province4 string
	var province5 string
	var province6 string
	var province7 string
	var province8 string
	var province9 string
	var province10 string
	var Territory1 string
	var Territory2 string
	var Territory3 string
	var cap1 string
	var cap2 string
	var cap3 string
	var cap4 string
	var cap5 string
	var cap6 string
	var cap7 string
	var cap8 string
	var cap9 string
	var cap10 string
	var cap11 string
	var cap12 string
	var cap13 string

//input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Can you name the Provinces of Canada? ")
reader.ReadString('\n')

	fmt.Print("What is the first one? ")
	province1, _ = reader.ReadString('\n')
	province1 = strings.TrimSpace(province1)

	fmt.Print("What is the second one? ")
	province2, _ = reader.ReadString('\n')
	province2 = strings.TrimSpace(province2)

	fmt.Print("What is the third one? ")
	province3, _ = reader.ReadString('\n')
	province3 = strings.TrimSpace(province3)

	fmt.Print("What is the fourth one? ")
	province4, _ = reader.ReadString('\n')
	province4 = strings.TrimSpace(province4)

	fmt.Print("What is the fifth one? ")
	province5, _ = reader.ReadString('\n')
	province5 = strings.TrimSpace(province5)

	fmt.Print("What is the sixth one? ")
	province6, _ = reader.ReadString('\n')
	province6 = strings.TrimSpace(province6)

	fmt.Print("What is the seventh one? ")
	province7, _ = reader.ReadString('\n')
	province7 = strings.TrimSpace(province7)

	fmt.Print("What is the eighth one? ")
	province8, _ = reader.ReadString('\n')
	province8 = strings.TrimSpace(province8)

	fmt.Print("What is the ninth one? ")
	province9, _ = reader.ReadString('\n')
	province9 = strings.TrimSpace(province9)

	fmt.Print("What is the tenth one? ")
	province10, _ = reader.ReadString('\n')
	province10 = strings.TrimSpace(province10)

	fmt.Println("Good Job!")

	fmt.Print("Now can you name the territories and every capital city? ")
	reader.ReadString('\n')

	fmt.Print("What is the first Territory? ")
	Territory1, _ = reader.ReadString('\n')
	Territory1 = strings.TrimSpace(Territory1)

	fmt.Print("What is the second Territory? ")
	Territory2, _ = reader.ReadString('\n')
	Territory2 = strings.TrimSpace(Territory2)

	fmt.Print("What is the last Territory? ")
	Territory3, _ = reader.ReadString('\n')
	Territory3 = strings.TrimSpace(Territory3)

	fmt.Println("Good Job!")

	fmt.Print("What is the first Capital you mentioned? ")
	cap1, _ = reader.ReadString('\n')
	cap1 = strings.TrimSpace(cap1)

	fmt.Print("What is the second Capital you mentioned? ")
	cap2, _ = reader.ReadString('\n')
	cap2 = strings.TrimSpace(cap2)

		fmt.Print("What is the third Capital you mentioned? ")
	cap3, _ = reader.ReadString('\n')
	cap3 = strings.TrimSpace(cap3)

	fmt.Print("What is the 4th Capital you mentioned? ")
	cap4, _ = reader.ReadString('\n')
	cap4 = strings.TrimSpace(cap4)

	fmt.Print("What is the 5th Capital you mentioned? ")
	cap5, _ = reader.ReadString('\n')
	cap5 = strings.TrimSpace(cap5)

	fmt.Print("What is the 6th Capital you mentioned? ")
	cap6, _ = reader.ReadString('\n')
	cap6 = strings.TrimSpace(cap6)

	fmt.Print("What is the 7th Capital you mentioned? ")
	cap7, _ = reader.ReadString('\n')
	cap7 = strings.TrimSpace(cap7)

	fmt.Print("What is the 8th Capital you mentioned? ")
	cap8, _ = reader.ReadString('\n')
	cap8 = strings.TrimSpace(cap8)

	fmt.Print("What is the 9th Capital you mentioned? ")
	cap9, _ = reader.ReadString('\n')
	cap9 = strings.TrimSpace(cap9)

	fmt.Print("What is the 10th Capital you mentioned? ")
	cap10, _ = reader.ReadString('\n')
	cap10 = strings.TrimSpace(cap10)

	fmt.Print("What is the 11th Capital you mentioned? ")
	cap11, _ = reader.ReadString('\n')
	cap11 = strings.TrimSpace(cap11)

	fmt.Print("What is the 12th Capital you mentioned? ")
	cap12, _ = reader.ReadString('\n')
	cap12 = strings.TrimSpace(cap12)

	fmt.Print("What is the 13th Capital you mentioned? ")
	cap13, _ = reader.ReadString('\n')
	cap13 = strings.TrimSpace(cap13)

//output
	fmt.Println("Thats very good!")

fmt.Println("Province/Territory                      Capital")
fmt.Println( province1 + "                             " + cap1)
fmt.Println( province2 + "                             " + cap2)
fmt.Println( province3 + "                             " + cap3)
fmt.Println( province4 + "                             " + cap4)
fmt.Println( province5 + "                             " + cap5)
fmt.Println( province6 + "                             " + cap6)
fmt.Println( province7 + "                             " + cap7)
fmt.Println( province8 + "                             " + cap8)
fmt.Println( province9 + "                             " + cap9)
fmt.Println( province10 + "                            " + cap10)
fmt.Println( Territory1 + "                            " + cap11)
fmt.Println( Territory2 + "                            " + cap12)
fmt.Println( Territory3 + "                            " + cap13)

	fmt.Println("\nDone")


}