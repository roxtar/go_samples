package main

import "runtime/pprof"
import "os"

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}

func main() {
  f, _ := os.Create("factorial.prof")
  defer f.Close()
  pprof.StartCPUProfile(f)
  defer pprof.StopCPUProfile()
  for i:=0; i < 20; i ++ {
    go println(factorial(i))
  }
}
