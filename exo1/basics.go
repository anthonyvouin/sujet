package exo1 

func Sum(a, b, c int) int {
	return a + b + c
}

func IsEven(a int) bool {
	return a%2 == 0
}

func MaxOfFour(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func CountOccurrences(s string, char rune) int {
	count := 0
	for _, r := range s {
		if r == char {
			count++
		}
	}
	return count
}

func FilterEven(numbers []int) []int {
	result := []int{}
	for _, n := range numbers {
		if IsEven(n) {
			result = append(result, n)
		}
	}
	return result
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
