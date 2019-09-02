# check_stack_driver.py はなんですか？
Nagios Pulugin作成トレーニングをするためのサンプルプログラムです。  
GCPのインスタンスの起動時間をチェックし、しきい値を超えて停止仕手いる場合アラートを上げます。  
  
## プラグインの使い方  
事前にGCPのSDKをインストールする  
`# pip install google-cloud-monitoring`  
  
コマンドの実行   
10分以上停止している場合、WARNINGとし、30分以上停止している場合、CRITICALとします。  
`# ./check_stack_driver.py --project "{project name}" --instance {instance name} --warning 10  --critical 30 --key "{key file path}"`  

## 引数
  
|引数|値|
|---|---|
|--project| GCPのプロジェクトID|
|--instance| GCEのインスタンスID|
|--critical| CRITICALしきい値（単位：分） |
|--warning| WARNING しきい値 （単位：分）|
|--key| StackDriverにアクセス可能なサービスアカウントのIAMキー|

## 言い訳
1. データ型のチェックが面倒だったのでしていません
2. もちろん未入力のチェックもない、空白を入れたときにどうなるかなんて知らない
  
これらは皆さんがより良いプラグインをすぐに作成するための、意図的な余白だ