package service

import (
	"log"
	"os"
	"testing"

	pb "github.com/Shahboz4131/to-do-service/genproto"

	"google.golang.org/grpc"
)

var client pb.TaskServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did  not connect %v", err)
	}

	client = pb.NewTaskServiceClient(conn)

	os.Exit(m.Run())
}
