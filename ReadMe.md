# Nagios Pluginの作り方

## ここを見るとスライドショーで見れます。
https://gitpitch.com/yuukichi-nankou/how-to-make-nagios-plugin

## Nagios プラグインのルール
* プラグインの Exit ステータス でサービスのステータスが決まる  
0 : OK  
1 : WARNING  
2 : CRITICAL  
3 : UNKNOWN  

* プラグインの標準出力がステータス情報として処理される。
* 出力に|（パイプ）が含まれている場合、パイプ以降はパフォーマンスデータとして処理される

### しきい値の指定  
-w : WARNING しきい値  
-c : CRITICAL しきい値  

|しきい値の指定|検知範囲|
|---|---|
|10|< 0 or > 10|
|10:|< 10|
|~:10|> 10|
|10:20|< 10 or > 20|
|@10:20|<= 10 or >= 20|

### ほかにも制約がある
デバッグオプションの指定方法など、詳細はこちら  
Nagios Plugins Development Guidelines  
http://nagios-plugins.org/doc/guidelines.html  