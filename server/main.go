package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
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

	// *** あくまでサンプルで用意した関数の内容 *** //

	dir := "/Users/sho/Coding/go/src/grpc/storage"

	paths, err := ioutil.ReadDir(dir) // ディレクトリ以下のリストを取得する = 今回だとstorage以下のPath = name.txt + occupation.txt
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

	// ループの作成
	for _, path := range paths {
		if !path.IsDir() { // パスがディレクトリでない場合 == ファイルである場合
			fileNames = append(fileNames, path.Name())
		}
	}

	// 戻り値のメッセージを作成するために初期化する
	res := &pb.ListFilesResponse{
		FileNames: fileNames,
	}

	return res, nil
}

// サーバー側の起動
func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Faild to listen %v", err)
	}

	s := grpc.NewServer()

	fmt.Println("Server is Running!")

	// grpc側に、作成したサーバーの構造体を登録 //
	pb.RegisterFileServiceServer(s, &server{})

	// 指定のリッスンポートでサーバーを起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Faild to serve %v", err)
	}

}

// *** go run server/main.go を行ってサーバー起動 *** //
