package main

import (
	"fmt"
	"math"
)

// Esempio di uso di type per definire **tipi** o **alias**
type Celsius int // trattato come "nuovo tipo"
type Kelvin  int // trattato come "nuovo tipo"
// type Celsius = int // trattato come "alias"
// type Kelvin  = int // trattato come "alias"

// Esempio di uso di type per definire **un nuovo tipo**
// struct (o record) con *campi* x, y, nome, v
type Punto struct {
	x float64
	y float64
	nome string
	v []int // un campo può avere tipo qualsiasi (anche composto)
}

// Esempio di operazione su elementi del nuovo tipo
// passaggio di parametri *sempre per copia*, campo per campo
// attenzione: p.v e` sempre per copia *della rappresentazione interna* (non del blocco di memoria puntato dalla slice)
func distanza( p Punto, q Punto ) float64 {

	return math.Sqrt( (p.x - q.x)*(p.x - q.x) + (p.y - q.y)*(p.y - q.y) )

}

func main() {
// sintassi struct "pura"
// var p struct{
//   x float64
//   y float64
//   nome string
//   v []int
// }

// Esempio: dichiarazione di variabili Punto, e uso di Punto in tipi composti
var p, q Punto
var m []Punto // un nuovo tipo può essere composto, come gli altri tipi di GO (es. slice di Punto)



	// Accesso in scrittura ai campi di una struct
	p.x = 12.2
	p.y = 51.7
	p.nome = "Punto A"

	// Letterali struct
	q = Punto{x: 21.2, y:4.5} // i campi che mancano sono valorizzati al **valore di inizializzazione default** per il loro tipo

	fmt.Println(q)

	// Esperimenti sulla rappresentazione interna di una struct

	q.x = 16.3
	fmt.Println(p, q)

	// assegnamento tra struct --> *copia* campo per campo
	q = p

	p.x = 9.0
	fmt.Println(p, q)

	fmt.Println( distanza(p, q) )

	// Struct e slice
	m = make([]Punto, 4)
	// m[0] si comporta in tutto e per tutto come una variabile di tipo Punto
	m[0] = p
	m[0].x = 88.8

	q = m[0]
	q.y = m[0].y

	// Un campo di una struct puo` anche essere di tipo composto
	m[0].v = make([]int, 10)
	
	fmt.Println(m)


}

