package main

import "fmt"

func main() {
	exemplo := []int{1, 2, 3, 4, 5}

	fmt.Println(soma(exemplo, 5, 7))
	fmt.Println(soma(exemplo, 5, 10))
}

func soma(v []int, n int, alvo int) (int, int) {
	i := 0
	j := n - 1

	for i < j {
		soma := v[i] + v[j]

		if soma == alvo {
			return v[i], v[j]
		} else if soma < alvo {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}