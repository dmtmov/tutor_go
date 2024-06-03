package main

import (
	// "encoding/binary"
	"fmt"
    "os/exec"
	"log"
	"strings"
)

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
	// cmd := exec.Command("tput", "cols")
	// cmd.Stdin = os.Stdin
	// out, err := cmd.Output()
	//    fmt.Printf("out: %v, cmd: %v", string(out), cmd)
	// if err != nil {
	// 	// log.Fatal(err)
	// 	panic(err)
	// }
	//
	// t = WindowDimensions{}
	// values := strings.Split(string(out), " ")

	// h := values[0]

	cmd := exec.Command("tput", "cols")
	cmd.Stdin = strings.NewReader("")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %v\n", out.String())
    // Need to discover why `tput cols/lines` is better than `stty size`
    // Maybe it's better to use 3d-party library for that.
	// THIS: https://github.com/atomicgo
	// examples: https://gobyexample.com/


	// height, err := strconv.Atoi(values[0])
	// width, err := strconv.Atoi(w)
	// t.height = height
	// t.width = width

	// middle := t.width % 2

	// fmt.Println(middle)

}
