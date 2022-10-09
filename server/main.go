package main

import "grpc/pb"

// file service serverの実装

type server struct {
	pb.UnimplementedFileServiceServer
}
