# gcsgcs
GCSからファイルを取得したりGCSにファイルをアップロードしたりするツールです

## Download File `dl`
バケット名`BucketName`とバケットをルートとしたファイルパス`GCSfilePath`を指定して実行します．ダウンロードしたファイルが`~/<BucketName>/<GCSfilePath>`に保存されます．
```
$ gcsgcs -do dl -b daiiz-bucket-1 -gf public/blank.png
```
