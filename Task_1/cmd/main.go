package main

// Подключаем необходимые пакеты
import (
	"Task_1/pkg/myFunc"
	"fmt"
	"os"
)

func main() {
	for {
		var num1, num2 float64
		var operator string

		// Запрашиваем у пользователя первое число
		for {
			fmt.Print("Введите первое число: ")
			_, err := fmt.Scan(&num1)
			if err == nil {
				break
			}
			// Выводим сообщение об ошибке, если введено некорректное число
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		}

		// Запрашиваем у пользователя операцию (+, -, *, /)
		for {
			fmt.Print("Выберите операцию (+, -, *, /): ")
			_, err := fmt.Scan(&operator)
			if err == nil && (operator == "+" || operator == "-" || operator == "*" || operator == "/") {
				break
			}
			// Выводим сообщение об ошибке, если введена некорректная операция
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		}

		// Запрашиваем у пользователя второе число
		for {
			fmt.Print("Введите второе число: ")
			_, err := fmt.Scan(&num2)
			if err == nil {
				break
			}
			// Выводим сообщение об ошибке, если введено некорректное число
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		}

		// Вычисляем результат с использованием функции из пакета myFunc
		result, err := myFunc.Calculate(num1, num2, operator)
		if err != nil {
			// Выводим сообщение об ошибке, если вычисление завершилось с ошибкой
			fmt.Println(err)
			return
		}

		// Выводим результат на экран
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

		// Запрос повторной операции
		var again string
		fmt.Print("\nВыполнить еще одну операцию? (да/нет): ")
		_, err = fmt.Scan(&again)
		if err != nil || again != "да" {
			break
		}
	}
	os.Exit(0)
}
