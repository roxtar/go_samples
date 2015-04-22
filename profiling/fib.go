package main

import "runtime/pprof"
import "os"

func fib(n int) int {
  if n <= 1 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}

func main() {
  f, _ := os.Create("fib.prof")
  defer f.Close()
  pprof.StartCPUProfile(f)
  defer pprof.StopCPUProfile()
  for i:=0; i < 45; i ++ {
    println("Iteration ", i)
    println(fib(i))
  }
}
