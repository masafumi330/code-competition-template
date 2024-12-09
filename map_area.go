package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Delimiter = "" // マップの区切り文字

const (
	dot   = "."
	ast   = "*"
	sharp = "#"
)

type Area [][]interface{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")
	H, _ := strconv.Atoi(line1[0]) // Height
	W, _ := strconv.Atoi(line1[1]) // Width

	var dotPoint = struct {
		X int
		Y int
	}{
		X: 0,
		Y: 0,
	}

	area := make(Area, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		row := make([]interface{}, W)
		line := strings.Split(scanner.Text(), Delimiter)
		for j := 0; j < W; j++ {
			row[j] = line[j]
			if line[j] == ast {
				dotPoint.X = i
				dotPoint.Y = j
			}
		}
		area[i] = row
	}

	// output
	// output := make(Area, H)
	for i := 0; i < H; i++ {
		// row := make([]interface{}, W)
		for j := 0; j < W; j++ {
			if area[i][j] == sharp {
				continue
			}
			if i == dotPoint.X && j == dotPoint.Y {
				area[i][j] = ast
			} else if i == dotPoint.X-1 && j == dotPoint.Y {
				area[i][j] = ast
			} else if i == dotPoint.X && j == dotPoint.Y-1 {
				area[i][j] = ast
			} else if i == dotPoint.X+1 && j == dotPoint.Y {
				area[i][j] = ast
			} else if i == dotPoint.X && j == dotPoint.Y+1 {
				area[i][j] = ast
			}
		}
		// output[i] = row
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(area[i][j])
		}
		fmt.Println()
	}
}
