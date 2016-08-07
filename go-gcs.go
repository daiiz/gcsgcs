package main

import (
    "fmt"
    "flag"
)

var (
    bucketName    = flag.String("b", "", "GCS Bucket name")
    filePath      = flag.String("gf", "", "GCS filePath or fileName")
    localFilePath = flag.String("lf", "", "Local filePath")
    doTask        = flag.String("do", "dl", "<dl | up>")
)

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
    }else if *doTask == "up" {
        if (len(*bucketName) > 0 && len(*filePath) > 0 && len(*localFilePath) > 0) {
            savedPath, err := Upload(*bucketName, *localFilePath, *filePath)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Println(savedPath)
        }
    }
}
