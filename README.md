# 定時スクレイピングツール

## 説明

- 社内用ツールで作成したツール
- プロダクト情報になるため、大きく簡易化したもの
  - Golangアウトプットのため

## 技術構成

- Go 1.12
- phantomjs
- goquery
- (MySQL)
- (Google Cloud Storage)

※括弧は社内用のみ

## 実稼働プロダクトの機能

- 定時（Mon 10:00 & Fri 23:00）に動く
- DBからユーザー登録済みのキーワード（複数）を取得
- ループを回して一件ずつスクレイピング
- エンコードしたキーワードを渡すとYahoo検索結果のHTMLとスクリーンショットを取得
- スクリーンショットはzipしCloud Storageへ
- HTMLは特定部分（リスティング広告）を取得
  - 設定URL
  - 表示URL
  - タイトル
  - 説明文
- 取得した上記タグはCSVにしてCloud Storageへ

## このコードについて

- 上記機能のうちスクリーンショットの保存及びタグ取得部分を抜粋したもの