package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 累積和: cumulative sum
// あらかじめ、配列の最初の値からの累積和を計算しておくことで、任意の区間の和をO(1)で求めることができる

// 一次元の累積和
// 日付とズレると混乱するから、揃えたほうがいいな

// example>
// input>
// 15 3
// 62 65 41 13 20 11 18 44 53 12 18 17 14 10 39
// 4 13
// 3 10
// 2 15

// output>
// 220
// 212
// 375

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(line1[0])
	Q, _ := strconv.Atoi(line1[1])

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	A := make([]int, N+1)
	for i, v := range line2 {
		A[i+1], _ = strconv.Atoi(v)
	}

	// 累積和を計算
	sums := make([]int, N+1)
	sums[0] = 0
	for i := 1; i <= N; i++ {
		sums[i] = sums[i-1] + A[i]
	}

	outputs := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])

		outputs[i] = sums[r] - sums[l-1]
	}
	for _, v := range outputs {
		fmt.Println(v)
	}
}
