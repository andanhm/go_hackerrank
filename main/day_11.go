package main

import (
	"fmt"
	"sort"
)

func main() {
	var row, col int = 6, 6
	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, col)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			var input int;
			fmt.Scan(&input);
			arr[i][j] = input
		}
	}
	sum := make([]int, 16)
	var h int = 0;
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum[h] = arr[i][j] + arr[i][j + 1] + arr[i][j + 2] + arr[i + 1][j + 1] + arr[i + 2][j] + arr[i + 2][j + 1] + arr[i + 2][j + 2];
			h++;
		}
	}
	sort.Ints(sum)
	fmt.Printf("%d", sum[15])
}