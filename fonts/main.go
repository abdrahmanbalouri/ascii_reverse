package main

import (
	"bufio"
	"fmt"
	"os"
)

var models map[byte][]string

func main() {
	models = BuildModels()

	
	file, err := os.Open("fonts/t.txt")
	if err != nil {
		fmt.Println("Can't find file : " )
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)

	var input []string

	  
		for scanner.Scan(){
			if scanner.Text()==""{
				input = append(input,"" )
			}
			
		input = append(input, scanner.Text())
	}
	fmt.Println(input)

	res := BuildReverseString(input)

	fmt.Println(res)
}

func PrintHelp() {
	fmt.Print("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
	os.Exit(0)
}

func BuildModels() map[byte][]string {
	artFile, err := os.Open("fonts/standard.txt")
	if err != nil {
		fmt.Println("Can't open art file.")
		fmt.Println(err)
		os.Exit(0)
	}

	lettersMap := make(map[byte][]string)
	scanner := bufio.NewScanner(artFile)

	currentLetter := ' '
	scanner.Scan()

	for scanner.Scan() {
		var model []string
		for i := 0; i < 8; i++ {
			model = append(model, scanner.Text())
			scanner.Scan()
		}
		lettersMap[byte(currentLetter)] = model
		currentLetter++

	}
	artFile.Close()
	return lettersMap
}

func BuildReverseString(input []string) string {
	if len(input[0]) == 0 {
		return ""
	}

	letter := ' '
       
	for letter <= '~' {
		flag := true
		letterSize := len(models[byte(letter)][0])
		
		for i := 0; i < 8; i++ {
			
			if len(input[0]) <letterSize || models[byte(letter)][i] != input[i][:letterSize] {
				flag = false
				break
			}
		}
		if flag {
			
			var nextInput []string
			for i := 0; i < 8; i++ {
				nextInput = append(nextInput, input[i][letterSize:])
				fmt.Println(len(nextInput[i]))
			}
			
			
			
			return string(letter) + BuildReverseString(nextInput)
		}

		letter++
	}
	return "This is not an ascii-art banner"
}

