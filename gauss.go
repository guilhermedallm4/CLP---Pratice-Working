   package main

   import (
      "fmt"
      "math/rand"
      "os"
      "strconv"
      "syscall"
      "time"
      "golang.org/x/sys/unix"
   )

   const MAXN int = 2000 // Max value of N
   var N int             // Matrix size

   // Matrices and vectors
   var A [MAXN][MAXN]float32
   var B, X [MAXN]float32 // A * X = B, solve for X

   func randm() int {
      var uid int
      rand.Seed(time.Now().UnixNano())
      return rand.Intn(4) | (2*uid)&3
   }

   func time_seed() uint32 {
      t := time.Now()
      return uint32(t.Nanosecond())
   }

   func parameters(argc int, argv []string) {
      seed := 0 // Random seed
      //var uid string // User name

      /* Read command-line arguments */
      var value_random = time_seed()
      rand.Seed(int64(value_random)) // Randomize

      if argc == 3 {
         seed, _ = strconv.Atoi(argv[2])
         rand.Seed(int64(seed))
         fmt.Printf("Random seed = %d\n", seed)
      }
      if argc >= 2 {
         N, _ = strconv.Atoi(argv[1])
         if N < 1 || N > MAXN {
            fmt.Printf("N = %d is out of range.\n", N)
            os.Exit(0)
         }
      } else {
         fmt.Printf("Usage: %s <matrix_dimension> [random seed]\n", argv[0])
         os.Exit(0)
      }

      /* Print parameters */
      fmt.Printf("\nMatrix dimension N = %d.\n", N)
   }

   func timeSeed() int64 {
      t := time.Now()
      return t.UnixNano() / int64(time.Millisecond)
   }

   func initialize_inputs() {
      fmt.Println("\nInitializing...")

      for col := 0; col < N; col++ {
         for row := 0; row < N; row++ {
            A[row][col] = float32(rand.Intn(32768)) / 32768.0
         }
         B[col] = float32(rand.Intn(32768)) / 32768.0
         X[col] = 0.0
      }
   }

   func print_X() {
      var row int

      if N < 100 {
         fmt.Printf("\nX = [")
         for row = 0; row < N; row++ {
            fmt.Printf("%5.2f%s", X[row], func() string {
               if row < N-1 {
                  return "; "
               }
               return "]\n"
            }())
         }
      }
   }

   func main() {
      // Timing variables
      var startTime, endTime syscall.Timeval
      var rusage unix.Rusage
      clkTck := float64(unix.Getpagesize())
      // get start time
      syscall.Gettimeofday(&startTime)
      // Process program parameters
      parameters(len(os.Args), os.Args)

      // Initialize A and B
      initialize_inputs()

      // Print input matrices
      initialize_inputs()

      // Gaussian Elimination
      gauss()

      //Stop program
      syscall.Gettimeofday(&endTime)
      // Display output
      print_X()

      // calculate elapsed time in milliseconds
      elapsedTime := float64(endTime.Sec)*1000 + float64(endTime.Usec/1000) - float64(startTime.Sec)*1000 - float64(startTime.Usec/1000)
      fmt.Printf("\nElapsed time = %g ms.\n", elapsedTime)

      // print accuracy of CPU times
      fmt.Printf("(CPU times are accurate to the nearest %g ms)\n", 1.0/clkTck*1000)

      // get CPU time for parent process
      unix.Getrusage(unix.RUSAGE_SELF, &rusage)
      cpuTime := float64(rusage.Utime.Sec)*1000 + float64(rusage.Utime.Usec/1000) + float64(rusage.Stime.Sec)*1000 + float64(rusage.Stime.Usec/1000)
      fmt.Printf("My total CPU time for parent = %g ms.\n", cpuTime/clkTck)

      // get system CPU time for parent process
      unix.Getrusage(unix.RUSAGE_SELF, &rusage)
      sysCpuTime := float64(rusage.Stime.Sec)*1000 + float64(rusage.Stime.Usec/1000)
      fmt.Printf("My system CPU time for parent = %g ms.\n", sysCpuTime/clkTck)

      // get CPU time for child processes
      unix.Getrusage(unix.RUSAGE_CHILDREN, &rusage)
      childCpuTime := float64(rusage.Utime.Sec)*1000 + float64(rusage.Utime.Usec/1000) + float64(rusage.Stime.Sec)*1000 + float64(rusage.Stime.Usec/1000)
      fmt.Printf("My total CPU time for child processes = %g ms.\n", childCpuTime/clkTck)
      /* Contrary to the man pages, this appears not to include the parent */
      fmt.Println("--------------------------------------------")

      os.Exit(0)
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
