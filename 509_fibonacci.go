//509. Fibonacci Number
//Runtime: 13 ms, faster than 20.19% of Go online submissions for Fibonacci Number.
//Memory Usage: 1.8 MB, less than 87.40% of Go online submissions for Fibonacci Number.
func fib(n int) int {
    switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return fib(n-1) + fib(n-2)
	}
}