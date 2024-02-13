# プロジェクト名

このプロジェクトは、ReactとReact Nativeを使用したフロントエンド開発、およびGo言語を使用したバックエンド開発を含むフルスタックアプリケーションです。

## 開始方法

### フロントエンド（React）

Reactプロジェクトを開始するには、以下の手順に従ってください。

1. プロジェクトディレクトリに移動します。
2. `npm start` を実行して、開発モードでアプリを起動します。
bash
cd React
npm start

詳細は [React/README.md](React/README.md) を参照してください。

### フロントエンド（React Native）

React Nativeプロジェクトを開始するには、以下の手順に従ってください。

1. プロジェクトディレクトリに移動します。
2. `npm start` を実行して、Metro Bundlerを起動します。
3. 新しいターミナルを開き、`npm run android` または `npm run ios` を実行して、アプリを起動します。
bash
cd ReactNativeApp
npm start

## 新しいターミナルで

npm run android # Androidの場合
npm run ios # iOSの場合

詳細は [ReactNativeApp/README.md](ReactNativeApp/README.md) を参照してください。

## バックエンド（Go）

Go言語で実装されたバックエンドを起動するには、以下の手順に従ってください。

1. `main.go` を実行してサーバーを起動します。
bash
go run main.go

2. GraphQL APIはデフォルトで `http://localhost:8080` で利用可能です。

詳細は [server.go](server.go) と [graph/generated.go](graph/generated.go) を参照してください。

## デプロイメント

アプリケーションのデプロイメントに関する情報は、各READMEファイルに記載されています。

- React: [React/README.md](React/README.md) のデプロイメントセクションを参照
- React Native: [ReactNativeApp/README.md](ReactNativeApp/README.md) のデプロイメントセクションを参照

## 貢献

このプロジェクトへの貢献に興味がある場合は、プルリクエストを送信するか、問題を報告してください。貢献ガイドラインについては、CONTRIBUTING.mdを参照してください。

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。詳細はLICENSEファイルを参照してください。
