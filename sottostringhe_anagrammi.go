package main

import (
	"fmt"
	"os"
)

func main() {
	m := map[string]int{}
	s := os.Args[1]
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if check1(s[i:j+1]) && check2(len(s[i:j+1])) {
				m[s[i:j+1]]++
			}
		}
	}

	for i := len(s); i > 0; i-- {
		for parola, occorenze := range m {
			if len(parola) == i {
				fmt.Printf("%v -> occorenze: %v\n", parola, occorenze)
			} else {
				continue
			}
		}
	}
}

func check1(x string) bool {
	if x[0] == x[len(x)-1] {
		return true
	}
	return false
}

func check2(lun int) bool {
	if lun >= 3 {
		return true
	}
	return false
}
