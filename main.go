package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	var result int

	fmt.Print("Введите выражение (например, 5+3 или V+III): ")
	fmt.Scan(&input)

	operatorPos := -1
	for i, char := range input {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			operatorPos = i
			break
		}
	}

	if operatorPos == -1 {
		fmt.Println("Неверный формат ввода. Пожалуйста, введите выражение в формате 'число+число'.")
		return
	}

	num1Str := input[:operatorPos]
	num2Str := input[operatorPos+1:]
	operator := input[operatorPos]

	var num1, num2 int
	var err1, err2 error

	if isRoman(num1Str) {
		num1 = romanToInt(num1Str)
	} else {
		num1, err1 = strconv.Atoi(num1Str)
	}

	if isRoman(num2Str) {
		num2 = romanToInt(num2Str)
	} else {
		num2, err2 = strconv.Atoi(num2Str)
	}

	if err1 != nil || err2 != nil {
		fmt.Println("Неверный формат чисел.")
		return
	}

	if num1 > 10 || num2 > 10 {
		fmt.Println("Неверный формат чисел. Числа не могут быть больше 10.")
		return
	}

	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Неподдерживаемый оператор.")
		return
	}

	if isRoman(num1Str) || isRoman(num2Str) {
		if result <= 0 {
			fmt.Println("Результат не может быть представлен в римских цифрах.")
		} else {
			fmt.Printf("Результат: %s\n", intToRoman(result))
		}
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}

func romanToInt(s string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	prevValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}
	return result
}

func isRoman(s string) bool {
	romanChars := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(romanChars, char) {
			return false
		}
	}
	return true
}