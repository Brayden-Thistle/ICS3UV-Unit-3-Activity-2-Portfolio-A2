/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-19
* @fileoverfiew this program sings the user happy birthday
*/

package main

import (
"fmt"
"strings"
"bufio"
"os"
)

func main () {
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
var agreement string


reader := bufio.NewReader(os.Stdin)

fmt.Print("Can you name the Provinces of Canada? ")
agreement, _ = reader.ReadString('\n')
agreement = strings.TrimSpace(agreement)

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
}