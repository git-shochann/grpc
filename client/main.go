package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"

	"google.golang.org/grpc"
)

// *** クライアントの実装(呼び出し元) 今回はGo同士で同士で通信してみる *** //

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // ターゲットのサーバーのアドレス + 接続方法 (練習のためこちらのメソッドで)
	if err != nil {
		log.Fatalf("Faild to connect %v", err)
	}
	defer connection.Close()

	// var client pb.FileServiceClient -> メソッドとして切り出す
	client := pb.NewFileServiceClient(connection)

	// ListFiles()をあとは呼ぶだけ

	callListFiles(client)

}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFileNames()) // ファイル名の一覧を出力！
}
