package main

import (
    "fmt"
    "flag"
)

var (
    bucketName = flag.String("b", "", "GCS Bucket name")
    filePath   = flag.String("gf", "", "GCS filePath")
    doTask     = flag.String("do", "dl", "<dl | up>")
)

// gcsgcs -do dl -b daiiz-bucket-1 -gf public/blank.png
func main() {
    flag.Parse()

    if *doTask == "dl" {
        if (len(*bucketName) > 0 && len(*filePath) > 0) {
            savedPath, err := Download(*bucketName, *filePath)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Println(savedPath)
        }
    }
}
