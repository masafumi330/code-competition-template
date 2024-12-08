package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Delimiter = "" // マップの区切り文字

type Area [][]interface{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")
	H, _ := strconv.Atoi(line1[0]) // Height
	W, _ := strconv.Atoi(line1[1]) // Width

	area := make(Area, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		row := make([]interface{}, W)
		line := strings.Split(scanner.Text(), Delimiter)
		for j := 0; j < W; j++ {
			row[j] = line[j]
		}
		area[i] = row
	}

	// output
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Printf("(%d, %d) = %v\n", i, j, area[i][j])
		}
	}
}
