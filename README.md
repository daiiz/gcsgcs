# gcsgcs
Google Cloud Storageからファイルを取得したりGCSにファイルをアップロードしたりするツールです

## 準備
* ローカルで実行する場合: gcloudツールを用いて，バケットを管理しているプロジェクトにgcloud auth loginしておく必要がある
* GCEのVMで実行する場合: バケットを管理しているプロジェクト内で起動しているVMであれば事前準備は必要ない

## Download File `dl`
バケット名`BucketName`とバケットをルートとしたファイルパス`GCSfilePath`を指定して実行します．ダウンロードしたファイルが`~/gcs_buckets/<BucketName>/<GCSfilePath>`に保存されます．
```
$ gcsgcs -do dl -b daiiz-bucket-1 -gf public/blank.png
```

## Upload File `up`
ローカルのファイルを，GCSのバケット名`BucketName`内のフォルダ`gcsgcs/`に`-gf`で指定したファイル名で保存します．
フォルダ`gcsgcs/`は予め作成しておいてください．
```
$ gcsgcs -do up -lf ./sample.txt -b daiiz-bucket-1 -gf sample.txt
```
