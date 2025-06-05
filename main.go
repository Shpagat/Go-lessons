package main

import "fmt"

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
	var first_num, second_num int
	// fmt.Println(romanToInt(text))

	fmt.Println("Как тебя зовут?") // Запрашиваем имя пользователя
	fmt.Scan(&text)
	fmt.Println("Привет,", text)

	fmt.Print("Введите 2 целых числа через пробел:") // Запрашиваем 2 числа
	fmt.Scan(&first_num, &second_num)

	fmt.Println("Сумма: ", first_num+second_num)        // Выводим сумму чисел
	fmt.Println("Разность: ", first_num-second_num)     // Выводим разность чисел
	fmt.Println("Произведение: ", first_num*second_num) // Выводим произведение чисел

	//	Проверяем не равно ли 2 число нулю
	if second_num != 0 {
		var first_num_dev float32 = float32(first_num)
		var second_num_dev float32 = float32(second_num)

		fmt.Println("Частное: ", (first_num_dev / second_num_dev))
	} else {
		fmt.Println("Делить на ноль нельзя")
	}
}
