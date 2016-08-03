package main

import (
    "os"
    "strings"
    "path/filepath"
    "bytes"
    "google.golang.org/cloud/storage"
    "golang.org/x/net/context"
    "github.com/mitchellh/go-homedir"
)

// 任意の場所に保存する
func SaveLocal(filePath string, content []byte) (string, error) {
    file, err := os.Create(filePath)
    if err != nil {
        return "", err
    }
    file.Write(content)
    defer file.Close()
    return filePath, nil
}

// ~/gcs_buckets/<bucket>/ にファイルを保存する
// ディレクトリ構造を維持する
func Download(bucket, filePath string) (string, error) {
    home, _ := homedir.Dir()

    dirs := strings.Split(filePath, "/")
    fileDir := ""
    fileName := dirs[len(dirs) - 1]
    for i := 0; i < len(dirs) - 1; i++ {
        fileDir += dirs[i] + "/"
    }

    dir := filepath.Join(home, "gcs_buckets")
    dir = filepath.Join(dir, bucket)
    dir = filepath.Join(dir, fileDir)

    // ローカルのホームにGCSのbucket階層構造を生成する
    if err := os.MkdirAll(dir, 0777); err != nil {
        return "", err
    }

    // GCSからファイル内容を取得する
    content, err := Fetch(bucket, filePath);
    if err != nil {
        return "", err
    }

    // ファイルをローカルに保存する
    savedFilePath, err := SaveLocal(filepath.Join(dir, fileName), content)
    if err != nil {
        return "", err
    }

    return savedFilePath, nil
}


func Fetch(bucket, filePath string) ([]byte, error) {
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
