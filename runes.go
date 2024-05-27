package main

import "fmt"
import "strings"

func RenderText(characters []Char) string {
	toPrint := make([]string, len(placeholder))
	for _, val := range characters {
		toPrint = append(toPrint, val.String())
	}

	return fmt.Sprintf("\r%v", strings.Join(toPrint, ""))

}
