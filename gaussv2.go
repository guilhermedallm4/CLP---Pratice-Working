package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAXN int = 2000 // Max value of N
var N int             // Matrix size

// Matrices and vectors
var A [MAXN][MAXN]float32
var B [MAXN]float32
var X [MAXN]float32

func parameters() {

	currentTime := time.Now()
	currentTimeUnix := int64(currentTime.Unix())
	args := os.Args
	rand.Seed(currentTimeUnix)
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
	if len(args) == 3 {
		currentTimeUnix, err := strconv.Atoi(args[2])
		if err == nil {
			rand.Seed(int64(currentTimeUnix))
			fmt.Printf("Random seed = %d\n", currentTimeUnix)
		}
	}
	if len(args) >= 2 {
		N, err := strconv.Atoi(args[1])
		if err != nil || N < 1 || N > MAXN {
			fmt.Printf("N = %s is out of range.\n", args[1])
			os.Exit(0)
		}
	} else {
		fmt.Printf("Usage: %s <matrix_dimension> [random seed]\n", args[0])
		fmt.Printf("N = %d is out of range.\n", N)
		os.Exit(0)
	}

	fmt.Printf("\nMatrix dimension N = %d.\n", N)
}

func initialize_inputs() {
	var row, col int

	fmt.Println("\nInitializing...")
	for col = 0; col < N; col++ {
		for row = 0; row < N; row++ {
			A[row][col] = float32(rand.Intn(100)) / 32768.0
		}
		B[col] = float32(rand.Intn(100)) / 32768.0
		X[col] = 0.0
	}
}

func print_inputs() {
	var row, col int

	if N < 10 {
		fmt.Println("\nA =")
		for row = 0; row < N; row++ {
			fmt.Printf("\t")
			for col = 0; col < N; col++ {
				fmt.Printf("%5.2f%s", A[row][col], func() string {
					if col < N-1 {
						return ", "
					}
					return ";\n\t"
				}())
			}
		}
		fmt.Println("\nB = [")
		for col = 0; col < N; col++ {
			fmt.Printf("%5.2f%s", B[col], func() string {
				if col < N-1 {
					return "; "
				}
				return "]\n"
			}())
		}
	}
}

func print_X() {
	if N < 100 {
		fmt.Printf("\nX = [")
		for row := 0; row < N; row++ {
			fmt.Printf("%5.2f%s", X[row], func() string {
				if row < N-1 {
					return "; "
				} else {
					return "]\n"
				}
			}())
		}
	}
}

func main() {
	/* Timing variables */
	var etstart, etstop time.Time /* Elapsed times using time.Now() */
	var etstart2, etstop2 float32 /* Elapsed times using time.Now().UnixNano() */

	/* Process program parameters */
	parameters()

	/* Initialize A and B */
	initialize_inputs()

	/* Print input matrices */
	print_inputs()

	/* Start Clock */
	fmt.Println("\nStarting clock.")
	etstart = time.Now()

	/* Gaussian Elimination */
	gauss()

	/* Stop Clock */
	etstop = time.Now()
	usecstart := etstart.UnixNano() / 1000
	usecstop := etstop.UnixNano() / 1000

	/* Display output */
	print_X()

	/* Display timing results */
	fmt.Printf("\nElapsed time = %g ms.\n", float32(usecstop-usecstart)/float32(1000))

	fmt.Printf("(CPU times are accurate to the nearest %g ms)\n", 1.0/float32(time.Second)*1000.0)
	fmt.Printf("My total CPU time for parent = %g ms.\n", float32((etstart2+etstop2))/float32(time.Second)*1000.0)
	fmt.Printf("My system CPU time for parent = %g ms.\n", float32(etstop2-etstart2)/float32(time.Second)*1000.0)
	fmt.Printf("My total CPU time for child processes = %g ms.\n", 0.0) // Não há processos filhos neste código
	fmt.Println("--------------------------------------------")
}

func gauss() {
	var norm, row, col int // Normalization row, and zeroing element row and col
	var multiplier float32
	fmt.Println("Computing Serially.")

	// Gaussian elimination
	for norm = 0; norm < N-1; norm++ {
		for row = norm + 1; row < N; row++ {
			multiplier = A[row][norm] / A[norm][norm]
			for col = norm; col < N; col++ {
				A[row][col] -= A[norm][col] * multiplier
			}
			B[row] -= B[norm] * multiplier
		}
	}

	// Back substitution
	for row = N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col = N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}
