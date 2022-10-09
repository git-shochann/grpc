package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"io/ioutil"
)

type server struct {
	pb.UnimplementedFileServiceServer // 構造体を埋め込む -> FileServiceServerインターフェースを満たしているのでメソッドを使用することが出来る
}

// *** sample_grpc.pb.go *** //

// type FileServiceServer interface {
// 	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
// 	mustEmbedUnimplementedFileServiceServer()
// }

// *** FileServiceServerインターフェースを実装している構造体 *** //
// type UnimplementedFileServiceServer struct {
// }

// func (UnimplementedFileServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
// }

// func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// *** 上書きする？ -> context.Contextでタイムアウトや認証関連で使用する *** //

func (s *server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("メソッドが呼び出されました")

	dir := "/Users/sho/Coding/go/src/grpc/storage"

	paths, err := ioutil.ReadDir(dir) // ストレージ配下のパスを取得する
	if err != nil {
		return nil, err
	}

	// *** 上記戻り値の型 ***
	// type FileInfo interface {
	// 	Name() string       // base name of the file
	// 	Size() int64        // length in bytes for regular files; system-dependent for others
	// 	Mode() FileMode     // file mode bits
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool        // abbreviation for Mode().IsDir()
	// 	Sys() interface{}   // underlying data source (can return nil)
	// }

	var fileNames []string // 格納する箱の用意

	// ループの作成 -> 何をしてる？
	for _, path := range paths {
		if !path.IsDir() { // パスがディレクトリでない場合 == ファイルでない場合
			fileNames = append(fileNames, path.Name())
		}
	}

	// 戻り値のメッセージを作成するために初期化する
	res := &pb.ListFilesResponse{
		FileNames: fileNames,
	}

	return res, nil

}
