package main

func fibonacciAt(f uint16) uint64 {
	return uint64(fibonacci(0, 1, f)) // 1.14 compiler will not tail for us
}

func fibonacci(c, t, p uint16) uint16 {
	if p == 0 {
		return c
	}
	return fibonacci(t, c+t, p-1)
}
