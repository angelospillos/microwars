package lib

// Fib do da fib dance
func Fib(f uint16) uint64 {
	return uint64(fibonacci(0, 1, f))
}

func fibonacci(c, t, p uint16) uint16 {
	if p == 0 {
		return c
	}
	return fibonacci(t, c+t, p-1)
}
