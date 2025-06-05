package main

import "fmt"

// import "bufio"

func romanToInt(s string) int {
	var number uint16 = 0
	var lenght int
	lenght = len(s)
	for i := lenght - 1; i >= 0; i-- {

		switch s[i] {

		case 'M':
			number += 1000
			if i == 0 {

			} else if s[i-1] == 'C' {
				number -= 100
				i--
			}
		case 'D':
			number += 500
			if i == 0 {

			} else if s[i-1] == 'C' {
				number -= 100
				i--
			}
		case 'C':
			number += 100
			if i == 0 {

			} else if s[i-1] == 'X' {
				number -= 10
				i--
			}
		case 'L':
			number += 50
			if i == 0 {

			} else if s[i-1] == 'X' {
				number -= 10
				i--
			}
		case 'X':
			number += 10
			if i == 0 {

			} else if s[i-1] == 'I' {
				number -= 1
				i--
			}
		case 'V':
			number += 5
			if i == 0 {

			} else if s[i-1] == 'I' {
				number -= 1
				i--
			}
		case 'I':
			number++
		}
	}
	return int(number)
}

func main() {

	var text string
	// fmt.Scan(&text)
	// fmt.Println(romanToInt(text))

	fmt.Println("Как тебя зовут?")
	fmt.Scan(&text)
	fmt.Println("Привет,", text)
}
