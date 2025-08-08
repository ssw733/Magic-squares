package main

import (
	"fmt"
	"sync"
	"time"
)

const dimension int = 4

// Computational Parallelism Depth (dimension = 3 -> depth = 1, dimension = 4 -> depth = 2, dimension = 5 -> depth = 3/4/5 ??)
const routineDepth int = 2

// Maximum number of square
const max int = dimension * dimension

var MagicConst int

// Array of numbers of square
var GenesisNumbers = [max]int{}

// Number of permutations of the original numbers of the square
var Factorial int = 1

// Counter of valid magic squares
var SquaresCounter int
var Mutex sync.Mutex
var Wg sync.WaitGroup

func init() {
	for i := range max {
		MagicConst += i + 1
		GenesisNumbers[i] = i + 1
		Factorial *= i + 1
	}
	MagicConst /= dimension
	fmt.Println("Dimension of square is", dimension)
	fmt.Println("Magic const is", MagicConst)
	fmt.Println("Factorial is", Factorial, "permutations")
	fmt.Println("Genesis numbers is", GenesisNumbers)
	fmt.Println("Starting ...")
}

func main() {
	start := time.Now()
	shuffleNumbers(GenesisNumbers, 0)
	Wg.Wait()
	fmt.Println("Total:", SquaresCounter, "squares")
	fmt.Println("Runtime is", time.Since(start))
}

func shuffleNumbers(square [max]int, pos int) {
	//pre validation of rows of sqaures
	if pos > 0 && pos%dimension == 0 {
		row := [dimension]int(square[pos-dimension : pos])
		if !checkRow(row) {
			return
		}
	}
	//main loop
	for i := pos; i < max; i++ {
		var square2 [max]int
		square2 = square
		square2[i], square2[pos] = square2[pos], square2[i]
		if pos <= routineDepth {
			Wg.Add(1)
			go func() {
				defer Wg.Done()
				shuffleNumbers(square2, pos+1)
			}()
		} else {
			shuffleNumbers(square2, pos+1)
		}
	}
	//check calculated squares
	if pos == max-1 && checkSquare(square) {
		printSquare(square)
	}
}

func checkRow(col [dimension]int) bool {
	var colSum int
	for _, v := range col {
		colSum += v
	}
	if colSum == MagicConst {
		return true
	} else {
		return false
	}
}

func checkSquare(square [max]int) bool {
	//check rows and cols
	for edge1 := range dimension {
		var rowSum, colSum int
		for edge2 := range dimension {
			rowSum += square[edge1*dimension+edge2]
			colSum += square[edge1+dimension*edge2]
		}
		if rowSum != MagicConst || colSum != MagicConst {
			return false
		}
	}
	//check diagonals
	var mainDiagSum, antiDiagSum int
	for diag := range dimension {
		mainDiagSum += square[dimension*diag+diag]
		antiDiagSum += square[(dimension-1)*(diag+1)]
	}
	if mainDiagSum != MagicConst || antiDiagSum != MagicConst {
		return false
	}
	Mutex.Lock()
	SquaresCounter++
	Mutex.Unlock()
	return true
}

func printSquare(square [max]int) {
	Mutex.Lock()
	/*magicSquare := [dimension][dimension]int{}
	for k, num := range square {
		magicSquare[int(math.Floor(float64(k))/float64(dimension))][k%dimension] = num
	}
	for _, line := range magicSquare {
		fmt.Println(line)
	}
	fmt.Println("-----")*/
	for k, v := range square {
		fmt.Print(v)
		if k < max-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("")
	Mutex.Unlock()
}
