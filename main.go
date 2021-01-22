package main
import (
"fmt"
"os"
"strings"
)

func main() {
  who := "Alice "
  if len(os.Args) > 1 {
    fmt.Println(strings.Join(os.Args[1:], "-"))
  }
  fmt.Println("Good Morning \n", who)
  fmt.Println("\n"+who[6:])
}