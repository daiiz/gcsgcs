package gcs

import (
    "os"
    "fmt"
    "bytes"
    "google.golang.org/cloud/storage"
    "golang.org/x/net/context"
)

func SaveToLocal(filePath string, content []byte) {
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Print(err)
        return
    }
    file.Write(content)
    defer file.Close()
}

func Get(bucket, filePath string) ([]byte, error) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    reader, err := client.Bucket(bucket).Object(filePath).NewReader(ctx)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    var buf bytes.Buffer
    if _, err := buf.ReadFrom(reader); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

func Foo () int {
    return 333
}
