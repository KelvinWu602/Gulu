package scanner

import (
	"fmt"

	"github.com/KelvinWu602/Gulu/core/token"
)

func Scan(s string) {
	x := token.Token{Type: token.AND, Lexeme: "weuif", Line: 123}
	fmt.Println(x)
}
