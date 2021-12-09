package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	ps := packstring(s)
	fmt.Print(ps)
}

func packstring(s string) string {

	var b strings.Builder
	var preisnumber bool
	var isintonly bool = true
	for k := 0; k < len(s); k++ {

		sk := string(s[k])
		i, err := strconv.Atoi(sk)
		if err == nil {
			if k == 0 || preisnumber == true {
				b.WriteString(sk)
				preisnumber = true
				continue
			}
			presk := string(s[k-1])
			for i > 1 {
				b.WriteString(presk)
				i--
			}
			preisnumber = true
		} else {
			b.WriteString(sk)
			preisnumber = false
			isintonly = false
		}

	}

	if isintonly == true {
		return "(некорректная строка)"
	} else {
		return b.String()
	}
}
