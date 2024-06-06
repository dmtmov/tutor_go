package main

import (
    "github.com/xyproto/ollamaclient"
)

const prompt = "provide some random code in golang in a row format as one line. provide just a code. no apostrophes at the edges"

func GetPlaceholderText() (Output string) {
    var oc = ollamaclient.NewWithModel("llama3:latest")

    // oc.Verbose = true
    if err := oc.PullIfNeeded(); err != nil {
        panic(err)
    }

    Output, err := oc.GetOutput(prompt)
    if err != nil {
        panic(err)
    }
    return
}
