package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	green        = "\033[32m"
	red          = "\033[31m"
	yellow       = "\033[33m"
	brightYellow = "\033[1;33m" // Bright Yellow for 'O'
	reset        = "\033[0m"
	clear        = "\033[H\033[2J"
)

func main() {
	height := 14
	tree := buildTree(height)
	displayBlinkingTree(tree, height)
}

func buildTree(height int) [][]string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	tree := make([][]string, height)
	for i := range tree {
		tree[i] = make([]string, 2*height)
		for j := range tree[i] {
			tree[i][j] = " "
		}
	}

	for i := 1; i <= height; i++ {
		for k := 1; k <= (2*i - 1); k++ {
			randNum := rand.Intn(10)
			position := height - i + k - 1 // Adjusted position index
			if randNum == 0 {
				tree[i-1][position] = brightYellow + "O" + reset
			} else if randNum < 4 {
				tree[i-1][position] = red + "X" + reset
			} else {
				tree[i-1][position] = green + "*" + reset
			}
		}
	}

	return tree
}

func displayBlinkingTree(tree [][]string, height int) {
	trunkWidth := height / 3
	trunkHeight := height / 6
	trunk := strings.Repeat("|", trunkWidth)

	printTree(tree, trunk, trunkWidth, trunkHeight) // Print the tree once before blinking starts

	blinkOn := true
	for {
		updateBlinking(tree, blinkOn)      // Update the blinking 'O's
		time.Sleep(500 * time.Millisecond) // Blink rate
		blinkOn = !blinkOn
	}
}

func printTree(tree [][]string, trunk string, trunkWidth, trunkHeight int) {
	for i, line := range tree {
		fmt.Print(strings.Repeat(" ", 14-i))
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}

	// Print the trunk
	for i := 0; i < trunkHeight; i++ {
		fmt.Print(strings.Repeat(" ", 14-trunkWidth/2))
		fmt.Print(trunk)
		fmt.Println()
	}
}

func updateBlinking(tree [][]string, blinkOn bool) {
	for i, line := range tree {
		for j, char := range line {
			if strings.Contains(char, "O") {
				if blinkOn {
					fmt.Printf("\033[%d;%dH%s", i+1, j+1, brightYellow+"O"+reset)
				} else {
					fmt.Printf("\033[%d;%dH ", i+1, j+1)
				}
			}
		}
	}
	fmt.Printf("\033[%d;0H", len(tree)+3) // Move cursor out of the tree area
}
