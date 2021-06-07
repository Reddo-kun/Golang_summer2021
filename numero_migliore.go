package main

import(
  "fmt"
  "os"
  "strconv"
  "math"
)
func main(){
  n := os.Args[1]
  d,_ := strconv.Atoi(os.Args[2])

  res := []string{n}

  for i := 0; i < d;  i++{
    res = rimuoviCifra(res)
  }

  //trovo il numero minore
  min := math.MaxInt32
  for _, i := range res {
    k,_ := strconv.Atoi(i)
    if k < min {
      min = k
    }
  }
  fmt.Println("numero migliore:", min)
}

func rimuoviCifra(list []string) []string {
  var res []string
  for _, n := range list {
    for i := range n {
      res = append(res, n[:i] + n[i+1:])
    }
  }
  fmt.Println(res)
  return res
}
