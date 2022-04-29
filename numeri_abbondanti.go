package main
import "fmt"

func main() {
  var soglia int
  fmt.Scan(&soglia)
  if soglia <= 0{
    fmt.Println("Soglia non valida")
  }else{
    NumeriAbbondanti(soglia)
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
func ÈAbbondante(n int) bool {
  return SommaDivisoriPropri(n) > n
}
func NumeriAbbondanti(limite int) {
  for i := 1; i < limite; i++{
    if ÈAbbondante(i) {fmt.Println(i)}
  }
}
