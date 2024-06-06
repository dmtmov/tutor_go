package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

// prints Words per minute at the end of script.
// `space` is a delimiter
func CalculateWPM() {
    fmt.Println("start")

    scanner := bufio.NewScanner(os.Stdin)
    startTime := time.Now()

    if scanner.Scan() {
        inputText := scanner.Text()
        endTime := time.Now()

        duration := endTime.Sub(startTime)
        words := strings.Fields(inputText)
        numWords := len(words)
        minutes := duration.Minutes()

        wpm := float64(numWords) / minutes
        fmt.Printf("typed %d words in %.2f mins. WPM: %.2f\n", numWords, minutes, wpm)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
}

