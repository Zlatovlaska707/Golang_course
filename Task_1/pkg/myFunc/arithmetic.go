package myFunc

import (
	"fmt"
)

// Calculate выполняет операцию между двумя числами в соответствии с заданной операцией.
// Возвращает результат вычисления и ошибку, если таковая возникла.
func Calculate(num1, num2 float64, operator string) (float64, error) {
	var result float64
	var err error

	switch operator {
	case "+":
		// Выполняем сложение
		result = num1 + num2
	case "-":
		// Выполняем вычитание
		result = num1 - num2
	case "*":
		// Выполняем умножение
		result = num1 * num2
	case "/":
		if num2 == 0 {
			// Проверка на деление на ноль
			err = fmt.Errorf("Ошибка: деление на ноль невозможно.")
		} else {
			// Выполняем деление
			result = num1 / num2
		}
	default:
		// Ошибка, если операция не поддерживается
		err = fmt.Errorf("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}

	return result, err
}
