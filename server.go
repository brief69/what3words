package main

import (
	"log"
	"my-what3words-clone/graph" // 独自のGraphQLスキーマを含むパッケージをインポート / Importing package containing custom GraphQL schema
	"net/http" // HTTPサーバー機能を提供するパッケージをインポート / Importing package providing HTTP server functionalities
	"os" // 環境変数へのアクセスを提供するパッケージをインポート / Importing package providing access to environment variables

	"github.com/99designs/gqlgen/graphql/handler" // GraphQLハンドラーを提供するパッケージをインポート / Importing package providing GraphQL handler
	"github.com/99designs/gqlgen/graphql/playground" // GraphQLプレイグラウンドハンドラーを提供するパッケージをインポート / Importing package providing GraphQL playground handler
)

func main() {
	
	const defaultPort = "8080" // デフォルトのポート番号を定義 / Defining default port number
	port := os.Getenv("PORT") // 環境変数からポート番号を取得 / Getting port number from environment variable
	if port == "" {
		port = defaultPort // 環境変数が設定されていない場合、デフォルトのポート番号を使用 / Using default port number if environment variable is not set
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})) // GraphQLサーバーを初期化 / Initializing GraphQL server

	http.Handle("/", playground.Handler("GraphQL playground", "/query")) // GraphQLプレイグラウンドのルートを設定 / Setting up route for GraphQL playground
	http.Handle("/query", srv) // GraphQLクエリのルートを設定 / Setting up route for GraphQL queries

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port) // サーバー起動メッセージをログに出力 / Logging server start message
	log.Fatal(http.ListenAndServe(":"+port, nil)) // サーバーを起動し、エラーが発生した場合はログに出力 / Starting server and logging any errors that occur
}
