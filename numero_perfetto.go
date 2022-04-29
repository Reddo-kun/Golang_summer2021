package main
import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  if ÈPerfetto(n){
    fmt.Println(n,"è perfetto")
  }else{
    fmt.Println(n,"non è perfetto")
  }
}
func SommaDivisoriPropri(n int) (somma int) {
  for i := 1; i < n; i++ {
    if n%i == 0{
      somma += i
    }
  }
  return 
}
func ÈPerfetto(n int) bool {
  return n > 0 && SommaDivisoriPropri(n) == n
}
