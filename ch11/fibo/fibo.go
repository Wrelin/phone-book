package fibo

// Fibonacci возвращает n-е число в последовательности Фибоначчи (начиная с 0-го: 0, 1, 1, 2, 3, ...).
// Для n < 0 возвращает -1 (ошибка).
func Fibonacci(n int) int {
	if n < 0 {
		return -1 // Обработка ошибки
	} else if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
