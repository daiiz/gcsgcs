# gcsgcs
Google Cloud Storageからファイルを取得したりGCSにファイルをアップロードしたりするツールです

## Download File `dl`
バケット名`BucketName`とバケットをルートとしたファイルパス`GCSfilePath`を指定して実行します．ダウンロードしたファイルが`~/gcs_buckets/<BucketName>/<GCSfilePath>`に保存されます．
```
$ gcsgcs -do dl -b daiiz-bucket-1 -gf public/blank.png
```

## Upload File `up` (Coming Soon!)
ローカルのファイルを，GCSのバケット名`BucketName`内のフォルダ`gcsgcs/`に`-gf`で指定したファイル名で保存します．
フォルダ`gcsgcs/`は予め作成しておいてください．
```
$ gcsgcs -do up -lf ./sample.txt -b daiiz-bucket-1 -gf sample.txt
```
