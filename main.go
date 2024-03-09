package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman_mode bool
var result string = ""

// проверка на содержание недопустимых символов входных данных
func valid_chars(str string) {
	for k := 0; k < len(str); k++ {
		if strings.IndexByte("0123456789+-*/IVXLCDM", str[k]) == -1 {
			panic(errors.New("введены некорректные символы"))
		}
	}
}

// проверяем, римские или арабские
func is_roman(str string) bool {
	for k := 0; k < len(str); k++ {
		if strings.IndexByte("IVXLCDM", str[k]) == -1 {
			return false
		}
	}
	return true
}

// определяем, какое именно римское число на входе
func what_roman(str string) (bool, int) {
	romans := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	j := 0
	f := false
	for i := 0; i < 10; i++ {
		if str == romans[i] {
			f = true
			j = i + 1
		}
	}
	return f, j
}

// перевод из арабского числа в римское
func arab_to_roman(ar int) string {
	//предполагается, что результат не больше 100, т.к. на входе два числа <=10
	l := len(strconv.Itoa(ar))
	znaki := strings.Split(strconv.Itoa(ar), "")
	roman_liters := [8]string{"", "I", "V", "X", "L", "C", "D", "M"}
	r := ""
	t := 0
	pos := 0
	for i := l - 1; i >= 0; i-- {
		k, _ := strconv.Atoi(znaki[l-i-1])
		t = (k + 1) / 5
		pos = (k + 1) % 5
		switch pos {
		case 0:
			{
				r = r + roman_liters[i*2+1] + roman_liters[i*2+t+1]
			}
		case 1:
			{
				r = r + roman_liters[t*(i*2+t+1)]
			}
		case 2:
			{
				r = r + roman_liters[t*(i*2+t+1)] + roman_liters[i*2+1]
			}
		case 3:
			{
				r = r + roman_liters[t*(i*2+t+1)] + roman_liters[i*2+1] + roman_liters[i*2+1]
			}
		case 4:
			{
				r = r + roman_liters[t*(i*2+t+1)] + roman_liters[i*2+1] + roman_liters[i*2+1] + roman_liters[i*2+1]
			}
		}

	}
	return r
}

func main() {

	var oper string = ""
	var roman_a, roman_b bool
	var a, b, c int = 0, 0, 0

	//ввод данных
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите операцию")
		text, _ := reader.ReadString('\n') //ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     //очищает все пустоты
		valid_chars(text)

		//ищем операцию
		oper_count := 0
		for i := 0; i < len(text); i++ {
			oper_count++
			switch text[i] {
			case '+':
				oper = "+"
			case '-':
				oper = "-"
			case '*':
				oper = "*"
			case '/':
				oper = "/"
			default:
				oper_count--
			}
		}

		// если нашли единственную операцию, то делим строку по ней
		if oper == "" {
			panic(errors.New("строка не является математической операцией"))
		} else if oper_count > 1 {
			panic(errors.New("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"))
		} else {

			numbers := strings.SplitN(text, oper, 2)

			//проверяем совпадение римских и арабских
			if (!is_roman(numbers[0]) && is_roman(numbers[1])) || (is_roman(numbers[0]) && !is_roman(numbers[1])) {
				panic(errors.New("используются одновременно разные системы счисления"))
			} else {
				//если оба римские
				if is_roman(numbers[0]) && is_roman(numbers[1]) {
					roman_mode = true
					roman_a, a = what_roman(numbers[0])
					roman_b, b = what_roman(numbers[1])
					if !roman_a || !roman_b {
						panic(errors.New("не удовлетворяет заданию — оба операнда должны быть целыми от 0 до 10"))
					}

				} else { //если оба арабские
					roman_mode = false
					a, _ = strconv.Atoi(numbers[0])
					b, _ = strconv.Atoi(numbers[1])
				}

			}

			if a > 10 || b > 10 {
				panic(errors.New("не удовлетворяет заданию — оба операнда должны быть <= 10"))
			} else {
				//вычисляем целый результат в зависимости от операции
				switch oper {
				case "+":
					c = a + b
				case "-":
					c = a - b
				case "*":
					c = a * b
				case "/":
					c = a / b
				}

				//Если Римские, то проверяем, что результат положительный
				if roman_mode {
					if c > 0 {
						//перевод результата в римские и вывод
						fmt.Println(arab_to_roman(c))
					} else {
						panic(errors.New("в римской системе нет отрицательных чисел"))
					}
				} else { //формируем вывод для арабских чисел
					fmt.Println(c)
				}

			}

		}

	}

}
