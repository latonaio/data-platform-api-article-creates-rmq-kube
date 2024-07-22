# data-platform-api-article-creates-rmq-kube
data-platform-api-article-creates-rmq-kube は、周辺システム　を データ連携基盤 と統合することを目的に、API で記事データを登録するマイクロサービスです。

* https://xxx.xxx.io/api/API_ARTICLE_SRV/creates/

## 動作環境
data-platform-api-article-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## 本レポジトリ が 対応する API サービス
data-platform-api-article-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス URL: https://xxx.xxx.io/api/API_ARTICLE_SRV/creates/

## 本レポジトリ に 含まれる API名
data-platform-api-article-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_Header（記事 - ヘッダ）
* A_Partner（記事 - 取引先）
* A_Address（記事 - 住所）
* A_Counter（記事 - カウンタ）
* A_Like（記事 - いいね）

## API への 値入力条件 の 初期値
data-platform-api-article-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール
Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMArticleCreates",
	"accepter": ["Header"],
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMArticleCreates",
	"accepter": ["All"],
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 記事 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"Article" ～ "IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
```
