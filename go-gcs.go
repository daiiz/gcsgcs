package main

import (
    "fmt"
)

func main() {
    savedPath, _ := Download("daiiz-bucket-1", "public/blank.png")
    fmt.Println(savedPath)
}
