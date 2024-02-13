package main

import (
	"github.com/graphql-go/graphql" // GraphQLのGo言語実装をインポートしています。/ Importing the Go implementation of GraphQL.
	"log" // ログ出力のためのパッケージをインポートしています。/ Importing the package for log output.
)

// Locationは緯度、経度、および単語を保持する構造体です。/ Location is a struct that holds latitude, longitude, and words.
type Location struct {
	Lat   float64 // 緯度 / Latitude
	Lng   float64 // 経度 / Longitude
	Words string  // 位置に関連付けられた単語 / Words associated with the location
}

// Wordsは単語を保持する構造体です。/ Words is a struct that holds words.
type Words struct {
	Words string // 単語 / Words
}

// locationTypeはGraphQLのLocation型を定義しています。/ Defining the GraphQL Location type.
var locationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Fields: graphql.Fields{
			"lat": &graphql.Field{
				Type: graphql.Float, // 緯度のフィールド定義 / Field definition for latitude
			},
			"lng": &graphql.Field{
				Type: graphql.Float, // 経度のフィールド定義 / Field definition for longitude
			},
			"words": &graphql.Field{
				Type: graphql.String, // 単語のフィールド定義 / Field definition for words
			},
		},
	})

// wordsTypeはGraphQLのWords型を定義しています。/ Defining the GraphQL Words type.
var wordsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Words",
	Fields: graphql.Fields{
		"words": &graphql.Field{
			Type: graphql.String, // 単語のフィールド定義 / Field definition for words
			},
		},
	})
	
// queryTypeはGraphQLのクエリ型を定義しています。/ Defining the GraphQL Query type.
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getLocationByWords": &graphql.Field{
			Type: locationType, // Location型を返す / Returns Location type
			Args: graphql.FieldConfigArgument{
				"words": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String), // 必須の文字列引数 / Required string argument
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// 単語から地点を検索するロジックを実装 / Implement logic to search location by words
				// 仮の実装として、固定の値を返します / As a placeholder, returning fixed values
				return &Location{
					Lat:   35.6895,
					Lng:   139.6917,
					Words: params.Args["words"].(string),
				}, nil
			},
		},
		"getWordsByLocation": &graphql.Field{
			Type: wordsType, // Words型を返す / Returns Words type
			Args: graphql.FieldConfigArgument{
				"lat": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float), // 必須の浮動小数点数引数（緯度）/ Required float argument (latitude)
				},
				"lng": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float), // 必須の浮動小数点数引数（経度）/ Required float argument (longitude)
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// 緯度と経度から最も近い地点の単語を検索するロジックを実装 / Implement logic to search words by closest location
				// 仮の実装として、固定の値を返します / As a placeholder, returning fixed values
				return &Words{
					Words: "example.word.here",
				}, nil
			},
		},
	},
}) // この行の終わりにセミコロンを追加しました。

// 既存のスキーマ定義を使用して新しいGraphQLスキーマを作成 / Creating a new GraphQL schema using existing schema definitions
var schema, err = graphql.NewSchema(graphql.SchemaConfig{Query: queryType}) // ':=' を 'var' と '=' に変更しました。
if err != nil {
	log.Fatalf("新しいGraphQLスキーマの作成に失敗しました。エラー: %v", err) // GraphQLスキーマの作成失敗時のエラーログ / Error log when failing to create GraphQL schema
}
_ = schema
// GraphQLサーバーのセットアップと実行... / Setting up and running the GraphQL server...