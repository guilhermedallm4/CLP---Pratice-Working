package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N int = 2000

// Matrices and vectors
var A [N][N]float32
var B, X [N]float32 // A * X = B, solve for X

func initialize_inputs() {
	fmt.Println("\nInitializing...")

	for col := 0; col < N; col++ {
		for row := 0; row < N; row++ {
			A[row][col] = float32(rand.Intn(32768)) / 32768.0
		}
		B[col] = 1.0
		X[col] = 0.0
	}
}

func main() {
	// Initialize A and B
	initialize_inputs()

	before := time.Now()

	// Gaussian Elimination
	gauss()

	elapsed := time.Since(before)

	var index int
	for index = 0; index < N; index++ {
		fmt.Printf("%f, ", X[index])
	}
	fmt.Println()
	fmt.Println("Time total : ", elapsed)
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
	// (Diagonal elements are not normalized to 1. This is treated in back substitution.)

	// Back substitution
	for row = N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col = N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}
