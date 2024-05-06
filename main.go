package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words = []string{
	"apple", "banana", "cherry", "date", "elderberry",
	"fig", "grape", "honeydew", "kiwi", "lemon",
	"mango", "nectarine", "orange", "pear", "quince",
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	scanner := bufio.NewScanner(os.Stdin)
	correctCount := 0
	timeLimit := 30 * time.Second

	fmt.Printf("タイピングゲームを開始します。制限時間は%v秒です。\n", timeLimit/time.Second)
	fmt.Println("Enter キーを押して開始してください。")
	scanner.Scan()

	timerCh := time.After(timeLimit)

gameLoop:
	for {
		word := words[r.Intn(len(words))]
		fmt.Printf("%s -> ", word)

		inputCh := make(chan string)
		go func() {
			scanner.Scan()
			inputCh <- scanner.Text()
		}()

		select {
		case <-timerCh:
			fmt.Println("\n\n制限時間が終了しました。")
			break gameLoop
		case input := <-inputCh:
			if strings.EqualFold(input, word) {
				correctCount++
			}
		}
	}

	fmt.Printf("%v秒間で%d問正解しました。\n", timeLimit/time.Second, correctCount)
}
