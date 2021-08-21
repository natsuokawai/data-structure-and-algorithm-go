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
	fmt.Println(join(a, " "))
	insertionSort(a)
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

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		for j := i - 1; j >= 0 && a[j] > key; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = key
		fmt.Println(join(a, " "))
	}
}
