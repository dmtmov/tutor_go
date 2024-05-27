package main

import (
	// "encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
) // "log"

func RenderText(characters []Char) string {
	toPrint := make([]string, len(placeholder))
	for _, val := range characters {
		toPrint = append(toPrint, val.String())
	}

	return fmt.Sprintf("\r%v", strings.Join(toPrint, ""))

}

type WindowDimensions struct {
	height int
	width  int
}

func (t WindowDimensions) getCurrent() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	// fmt.Printf("out: %v\n", string(out))
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}

	t = WindowDimensions{}
	values := strings.Split(string(out), " ")

	// h := values[0]
	w := strings.TrimSuffix(values[1], "\n")

    // Need to discover why `tput cols/lines` is better than `stty size`
    // Maybe it's better to use 3d-party library for that.
	// THIS: https://github.com/atomicgo
	// examples: https://gobyexample.com/

	height, err := strconv.Atoi(values[0])
	width, err := strconv.Atoi(w)
	t.height = height
	t.width = width

	fmt.Println(t)

	middle := t.width % 2

	fmt.Println(middle)

}
