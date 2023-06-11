package fib

// Fib returns nth Fibonnaci sequence number.
func Fib(n uint) uint {
	if n <= 1 {
		return 1
	}
	return Fib(n-2) + Fib(n-1)
}
