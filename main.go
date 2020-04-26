package main

import (
    "fmt"
    "log"

    phoneinfoga "gopkg.in/sundowndev/phoneinfoga.v2/pkg/scanners"
)

func main() {
    number, err := phoneinfoga.LocalScan("<number>")

    if err != nil {
        log.Fatal(err)
    }

    links := phoneinfoga.GoogleSearchScan(number)

    for _, link := range links.Individuals {
        fmt.Println(link.URL) // Google search link to scan
    }
}

