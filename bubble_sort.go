package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n, _ := strconv.Atoi(read())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(read())
	}
	swap := bubbleSort(a)
	fmt.Println(join(a, " "))
	fmt.Println(swap)
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func join(a []int, sep string) string {
	var str, _sep string
	for _, n := range a {
		str += fmt.Sprintf("%s%s", _sep, strconv.Itoa(n))
		_sep = sep
	}
	return str
}

func bubbleSort(a []int) int {
	var swap int
	for i := 0; i < len(a)-1; i++ {
		for j := len(a) - 1; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				swap++
			}
		}
	}
	return swap
}
