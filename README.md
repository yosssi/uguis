# Uguis - Twitter音声読み上げツール

## 概要

Uguisは、Twitterのツイートを自動的に取得して音声で読み上げるコマンドラインツールです。[VoiceText Web API](https://cloud.voicetext.jp/webapi)を利用してテキスト情報（ツイート）を音声ファイルに変換します。

## インストール

### 方法1. 実行ファイルのインストール

準備中

### 方法2. ソースコードからのインストール

```sh
$ go get github.com/yosssi/uguis/...
```

## インストールが正常に実施されたことの確認

```sh
$ uguis -v
Uguis vX.X.X
```

## 起動に必要な認証情報の取得

以下APIの認証情報を取得してください。

* [Twitter REST API](https://dev.twitter.com/)

取得する認証情報：APIキー、APIシークレット、アクセストークン、アクセストークンシークレット

* [VoiceText Web API](https://cloud.voicetext.jp/webapi)

取得する認証情報：APIキー

## 起動方法

起動時に入力するコマンドは以下の通りです。

```sh
$ uguis -twitter-api-key=TwitterAPIキー -twitter-api-secret=TwitterAPIシークレット -twitter-access-token=Twitterアクセストークン -twitter-access-token-secret=Twitterアクセストークンシークレット -voicetext-api-key=VoiceTextAPIキー -p=音声ファイル再生コマンド
```

コマンド入力例を以下に記載します。

```sh
uguis -twitter-api-key=aaaa -twitter-api-secret=bbbb -twitter-access-token=cccc -twitter-access-token-secret=dddd -voicetext-api-key=eeee -p=afplay
```

## 取得・読み上げることのできるツイート

* 自分のタイムラインに表示されるツイート
* 特定のキーワードに合致するツイート（準備中）

