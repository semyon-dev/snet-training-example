package main

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	servicePB "snet-training-example/service"
	pb "snet-training-example/training"
	"strconv"
)

var (
	serverAddr = "0.0.0.0:5001"
)

type model struct {
	name   string
	method string
	desc   string
	status string
}

var models = map[string]model{}

// protoc -I . main.proto --go-grpc_out=. --go_out=.
func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server started")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	var trainingServer TrainServer
	pb.RegisterModelServer(grpcServer, &trainingServer)
	servicePB.RegisterProMethodsServer(grpcServer, &trainingServer)
	servicePB.RegisterBasicMethodsServer(grpcServer, &trainingServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

type TrainServer struct {
	pb.UnimplementedModelServer
	servicePB.UnimplementedProMethodsServer
	servicePB.UnimplementedBasicMethodsServer
}

func (s *TrainServer) Stt(context.Context, *servicePB.SttInput) (*servicePB.SttResp, error) {
	log.Println("Stt request")
	return &servicePB.SttResp{
		Result: "success stt (mock reply)",
	}, nil
}

func (s *TrainServer) Tts(context.Context, *servicePB.TtsInput) (*servicePB.TtsResponse, error) {
	log.Println("Tts request")
	return &servicePB.TtsResponse{
		Text: "success tts (mock reply)",
	}, nil
}

func (s *TrainServer) CreateModel(ctx context.Context, newModel *pb.NewModel) (*pb.ModelID, error) {
	// Generate a random integer
	randomIntInRange := rand.Intn(300) // Generates a random integer between 0 and 100 (inclusive)
	strModelID := strconv.Itoa(randomIntInRange)
	fmt.Println("new model, id:", strModelID)
	models[strModelID] = model{
		name:   newModel.Name,
		method: newModel.GrpcMethodName,
		desc:   newModel.Description,
	}
	return &pb.ModelID{
		ModelId: strModelID,
	}, nil
}

func (s *TrainServer) ValidateModelPrice(ctx context.Context, request *pb.ValidateRequest) (*pb.PriceInBaseUnit, error) {
	log.Println("ValidateModelPrice request")
	return &pb.PriceInBaseUnit{
		Price: 1,
	}, nil
}

func (s *TrainServer) UploadAndValidate(stream pb.Model_UploadAndValidateServer) error {
	log.Println("UploadAndValidate started")
	var fullData bytes.Buffer // Для хранения всего принятого файла
	var modelID string
	for {
		// Получаем сообщение из потока
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("catched EOF")
			// Клиент завершил отправку данных
			break
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		log.Printf("Received chunk of data for model %s", modelID)
		fullData.Write(req.Data) // Добавляем данные в общий буфер
	}
	err := os.WriteFile("test.txt", fullData.Bytes(), 0644)
	if err != nil {
		log.Println("can't write file: ", err)
	}
	log.Printf("Received file for model %s with size %d bytes", modelID, fullData.Len())
	return stream.SendAndClose(&pb.StatusResponse{
		Status: pb.Status_VALIDATED,
	})
}

func (s *TrainServer) ValidateModel(ctx context.Context, request *pb.ValidateRequest) (*pb.StatusResponse, error) {
	log.Println("validate model")
	return &pb.StatusResponse{
		Status: pb.Status_VALIDATING,
	}, nil
}

func (s *TrainServer) TrainModelPrice(ctx context.Context, id *pb.ModelID) (*pb.PriceInBaseUnit, error) {
	log.Println("train model price")
	return &pb.PriceInBaseUnit{
		Price: 1,
	}, nil
}

func (s *TrainServer) TrainModel(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("TrainModel request")
	return &pb.StatusResponse{
		Status: pb.Status_TRAINING,
	}, nil
}

func (s *TrainServer) DeleteModel(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("delete request")
	return &pb.StatusResponse{
		Status: pb.Status_DELETED,
	}, nil
}

func (s *TrainServer) GetModelStatus(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("GetModelStatus request")
	return &pb.StatusResponse{
		Status: pb.Status_VALIDATING,
	}, nil
}

func (s *TrainServer) mustEmbedUnimplementedModelServer() {
	panic("implement me")
}
