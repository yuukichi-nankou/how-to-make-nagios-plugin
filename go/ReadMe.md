# check_stack_driver.go はなんですか？
Nagios Pulugin作成トレーニングをするためのサンプルプログラムです。  
GCPのインスタンスの起動時間をチェックし、しきい値を超えて停止している場合アラートを上げます。  
  
## プラグインの使い方 
事前にGCPのSDKを取得して、プラグインをビルドする
```
# go get -u cloud.google.com/go/monitoring/apiv3
# go get -u github.com/jessevdk/go-flags
# go build check_stack_driver.go
```

コマンドの実行   
10分以上停止している場合、WARNINGとし、30分以上停止している場合、CRITICALとします。  
`# ./check_stack_driver --project "{project name}" --instance {instance name} --warning 10  --critical 30 --auth "{key file path}"`  

## 引数
  
|引数|値|
|---|---|
|--project| GCPのプロジェクトID|
|--instance| GCEのインスタンスID|
|--critical| CRITICALしきい値（単位：分） |
|--warning| WARNING しきい値 （単位：分）|
|--auth| StackDriverにアクセス可能なサービスアカウントのIAMキー|