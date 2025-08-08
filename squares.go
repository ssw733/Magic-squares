package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

const dimension int = 4

const routineDepth int = 2

// Maximum number of square
const max int = dimension * dimension

// Magic const and Counter of valid magic squares
var MagicConst, SquaresCounter int

// Array of numbers of square
var GenesisNumbers = [max]int{}

// Number of permutations of the original numbers of the square
var Factorial int = 1

var MagicRows map[[dimension]int][][dimension]int
var HarmonicMagicRowsKeys [][dimension]int
var Mutex sync.Mutex
var Wg sync.WaitGroup
var Lol int

func init() {
	for i := range max {
		MagicConst += i + 1
		GenesisNumbers[i] = i + 1
		Factorial *= i + 1
	}
	MagicConst /= dimension
	MagicRows = make(map[[dimension]int][][dimension]int)
	fmt.Println("Dimension of square is", dimension)
	fmt.Println("Magic const is", MagicConst)
	fmt.Println("Factorial is", Factorial, "permutations")
	fmt.Println("Genesis numbers is", GenesisNumbers)
	fmt.Println("Starting ...")
}

func main() {
	start := time.Now()
	findRows([dimension]int{}, 0)
	Wg.Add(1)
	go func() {
		defer Wg.Done()
		findHarmonicRows([max]int{}, 0)
	}()
	Wg.Wait()
	fmt.Println("Total:", SquaresCounter, "squares")
	fmt.Println("Runtime is", time.Since(start))
}

func findRows(row [dimension]int, pos int) {
	for i := 0; i < max; i++ {
		if !inRow(row, i+1) {
			var row2 [dimension]int
			row2 = row
			if pos < dimension {
				row2[pos] = i + 1
			}
			if pos < dimension-1 {
				findRows(row2, pos+1)
			}
			if pos == dimension-1 && checkRow(row2) {
				rowPseudonim := findRowPseudonim(row2)
				Mutex.Lock()
				MagicRows[rowPseudonim] = append(MagicRows[rowPseudonim], row2)
				Mutex.Unlock()
			}
		}
	}
}

func findHarmonicRows(square [max]int, pos int) {
	for key, _ := range MagicRows {
		if pos == 0 || pos > 0 && !numbersOfRowInSquare(square, key, pos*dimension) {
			var square2 [max]int
			square2 = square
			if pos < dimension {
				for k, v := range key {
					square2[dimension*pos+k] = v
				}
			}
			if pos <= routineDepth {
				Wg.Add(1)
				go func() {
					defer Wg.Done()
					findHarmonicRows(square2, pos+1)
				}()
			} else if pos < dimension-1 {
				findHarmonicRows(square2, pos+1)
			}
			if pos == dimension-1 {
				shuffleRows([max]int{}, square2, 0)
			}
		}
	}
}

func shuffleRows(square [max]int, squarePattern [max]int, pos int) {
	var key [dimension]int
	for i := 0; i < dimension; i++ {
		key[i] = squarePattern[pos*dimension+i]
	}
	for _, row := range MagicRows[key] {
		var square2 [max]int
		square2 = square
		for i, n := range row {
			square2[pos*dimension+i] = n
		}
		if pos < dimension-1 {
			shuffleRows(square2, squarePattern, pos+1)
		} else if pos == dimension-1 && checkSquare(square2) {
			printSquare(square2)
		}
	}
}

func findRowPseudonim(row [dimension]int) [dimension]int {
	sort.Ints(row[:])
	return row
}

func numbersOfRowInSquare(square [max]int, row [dimension]int, pos int) bool {
	for i := 0; i < pos; i++ {
		for _, v := range row {
			if square[i] == v {
				return true
			}
		}

	}
	return false
}

func inRow(row [dimension]int, number int) bool {
	for _, v := range row {
		if v == number {
			return true
		}
	}
	return false
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
	for i1, v1 := range square {
		for i2, v2 := range square {
			if i1 != i2 && v1 == v2 {
				return false
			}
		}
	}
	Mutex.Lock()
	SquaresCounter++
	Mutex.Unlock()
	return true
}

func printSquare(square [max]int) {
	Mutex.Lock()
	magicSquare := [dimension][dimension]int{}
	for k, num := range square {
		magicSquare[int(math.Floor(float64(k))/float64(dimension))][k%dimension] = num
	}
	for _, line := range magicSquare {
		fmt.Println(line)
	}
	fmt.Println("-----")
	//fmt.Print(strings.Trim(fmt.Sprint(square), "[]") + "\n")
	Mutex.Unlock()
}
