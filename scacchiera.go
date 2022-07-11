// Un semplice simulatore per giocare a scacchi

package main

import "fmt"

// La funzione main simula una partita a scacchi tra due giocatori.
// Usa le funzioni definite in scacchiera.go.
func main() {
	fmt.Println("semplice simulatore per gli scacchi\nsintassi mossa:\nriga1 col1 riga2 col2")
	fmt.Println("(riga compreso tra 1..8; col tra 'a'..'h')")
	fmt.Println("Ctrl-c per terminare in anticipo")
	fmt.Println("impostare il terminale utf-8 e con sfondo bianco\n")
	// inizializziamo la scacchiera, la visualizziamo
	s := InizializzaScacchiera()
	PrintScacchiera(s)
	// ... e iniziamo a giocare
	prima := map[Colore]bool{BIANCO: true, NERO: true} // prima mossa per ciascun colore
	namecol := map[Colore]string{BIANCO: "Bianco", NERO: "Nero"}
	turno := Colore(BIANCO) //inizia il bianco
	for mangiato := Tipo(NULL); mangiato != RE; {
		fmt.Printf("muove il %s: ", namecol[turno])
		c1, c2 := LeggiMossa()
		if src, dest, ok := Muovi(s, turno, prima[turno], c1, c2); ok {
			fmt.Printf("mosso:%s mangiato:%s\n", StringPezzo(src), StringPezzo(dest))
			mangiato = dest.tipo
			prima[turno] = false
			turno = !turno
		} else {
			fmt.Println("mossa non valida")
		}
		PrintScacchiera(s)
	}
	fmt.Printf("scacco! vince il %s\n", namecol[!turno])
}

// PrintScacchiera visualizza il contenuto corrente di una scacchiera su stdout.
// ATTENZIONE! Questa versione assume che Scacchiera sia implementato come map.
func PrintScacchiera(s Scacchiera) {
	fmt.Printf("%4s", "\u2502")
	for c := 'a'; c < 'a'+8; c++ { // riga 0 con le lettere per denotare le col.
		fmt.Printf("%3c  ", c)
	}
	hline := hline(5) // linea orizzontale
	fmt.Println(hline)
	var r, c byte
	for r = 8; r != 0; r-- {
		fmt.Printf("%-3d\u2502", r)
		for c = 'a'; c < 'a'+8; c++ {
			t := StringPezzo(s[Casella{r, c}])
			fmt.Printf("%2s  \u2502", t)
		}
		fmt.Println(hline)
	}
}

// hline ritorna una linea orizzontale in cui ogni elemento ha dimensione width
func hline(width int) string {
	r := "\n"
	for i := 0; i < 9*width; i++ {
		r += "\u2500"
	}
	return r
}
