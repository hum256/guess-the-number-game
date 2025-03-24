package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// readInts は scanner から expected 個の整数を読み込み、整数のスライスを返します。
func readInts(scanner *bufio.Scanner, expected int) ([]int, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("入力読み込みに失敗しました。")
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) != expected {
		return nil, fmt.Errorf("整数%d個の入力が必要です。", expected)
	}
	nums := make([]int, expected)
	for i, s := range fields {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("数当てゲームへようこそ！")

	var lower, upper int
	// 下限と上限の入力ループ
	for {
		fmt.Print("下限と上限をスペース区切りで入力してください：")
		nums, err := readInts(scanner, 2)
		if err != nil {
			fmt.Println(err)
			continue
		} else if nums[0] >= nums[1] {
			fmt.Println("下限、上限の順に入力してください")
			continue
		}
		lower, upper = nums[0], nums[1]
		break
	}

	answer := rand.Intn(upper-lower+1) + lower
	// 数当ての入力ループ
	for {
		fmt.Print("数を予想してください：")
		nums, err := readInts(scanner, 1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		guess := nums[0]
		if guess < answer {
			fmt.Println("もっと大きいです。")
		} else if guess > answer {
			fmt.Println("もっと小さいです。")
		} else {
			fmt.Println("正解です！")
			break
		}
	}
}
