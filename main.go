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
	pb "snet-training-example/service"
	"strconv"
	"time"
)

var (
	serverAddr = "0.0.0.0:5001"
)

type model struct {
	name   string
	method string
	Status pb.Status `json:"status"`
}

var models = map[string]model{}

// protoc -I . service.proto --go-grpc_out=. --go_out=.
func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server started on " + serverAddr)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	var proMethods ExampleServer
	pb.RegisterModelServer(grpcServer, &proMethods)
	pb.RegisterExampleServiceServer(grpcServer, &proMethods)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

type ExampleServer struct {
	pb.UnimplementedExampleServiceServer
	pb.UnimplementedModelServer
}

func (s *ExampleServer) mustEmbedUnimplementedModelServer() {
	panic("implement me")
}

func (s *ExampleServer) BasicStt(c context.Context, r *pb.BasicSttInput) (*pb.SttResp, error) {
	log.Println("basic stt request")
	return &pb.SttResp{Result: "you are using service without model id"}, nil
}

func (s *ExampleServer) Stt(c context.Context, r *pb.SttInput) (*pb.SttResp, error) {
	return &pb.SttResp{Result: "you are using service with modelID" + r.ModelId.ModelId}, nil
}

func (s *ExampleServer) CreateModel(ctx context.Context, newModel *pb.NewModel) (*pb.ModelID, error) {
	// Generate a random integer
	randomIntInRange := rand.Intn(100000) // Generates a random integer between 0 and 100000 (inclusive)
	strModelID := strconv.Itoa(randomIntInRange)
	fmt.Println("new model, id:", strModelID)
	models[strModelID] = model{
		name:   newModel.Name,
		method: newModel.GrpcMethodName,
		Status: pb.Status_CREATED,
	}
	return &pb.ModelID{
		ModelId: strModelID,
	}, nil
}

func (s *ExampleServer) ValidateModelPrice(ctx context.Context, request *pb.ValidateRequest) (*pb.PriceInBaseUnit, error) {
	log.Println("ValidateModelPrice request")
	return &pb.PriceInBaseUnit{
		Price: 1,
	}, nil
}

func (s *ExampleServer) UploadAndValidate(stream pb.Model_UploadAndValidateServer) error {
	log.Println("UploadAndValidate started")
	var fullData bytes.Buffer // Для хранения всего принятого файла
	var modelID string
	var name string
	for {
		// Получаем сообщение из потока
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("catched EOF")
			// Клиент завершил отправку данных
			break
		}
		if req == nil {
			continue
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			break
		}
		name = req.FileName
		modelID = req.ModelId
		log.Printf("Received chunk of data for model %s", modelID)
		fullData.Write(req.Data) // Добавляем данные в общий буфер
	}
	err := os.WriteFile(name, fullData.Bytes(), 0644)
	if err != nil {
		log.Println("can't write file: ", err)
	}
	log.Printf("Received file for model %s with size %d bytes", modelID, fullData.Len())
	go func() {
		models[modelID] = model{Status: pb.Status_VALIDATING}
		time.Sleep(3 * time.Second)
		models[modelID] = model{Status: pb.Status_VALIDATED}
	}()
	return stream.SendAndClose(&pb.StatusResponse{
		Status: pb.Status_VALIDATING,
	})
}

func (s *ExampleServer) ValidateModel(ctx context.Context, req *pb.ValidateRequest) (*pb.StatusResponse, error) {
	log.Println("validate model")
	go func() {
		models[req.ModelId] = model{Status: pb.Status_VALIDATING}
		time.Sleep(3 * time.Second)
		models[req.ModelId] = model{Status: pb.Status_VALIDATED}
	}()
	return &pb.StatusResponse{
		Status: pb.Status_VALIDATING,
	}, nil
}

func (s *ExampleServer) TrainModelPrice(ctx context.Context, id *pb.ModelID) (*pb.PriceInBaseUnit, error) {
	log.Println("train model price")
	return &pb.PriceInBaseUnit{
		Price: 1,
	}, nil
}

func (s *ExampleServer) TrainModel(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("TrainModel request")
	go func() {
		models[id.ModelId] = model{Status: pb.Status_TRAINING}
		time.Sleep(3 * time.Second)
		models[id.ModelId] = model{Status: pb.Status_READY_TO_USE}
	}()
	return &pb.StatusResponse{
		Status: pb.Status_TRAINING,
	}, nil
}

func (s *ExampleServer) DeleteModel(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("delete request")
	models[id.ModelId] = model{Status: pb.Status_DELETED}
	return &pb.StatusResponse{
		Status: pb.Status_DELETED,
	}, nil
}

func (s *ExampleServer) GetModelStatus(ctx context.Context, id *pb.ModelID) (*pb.StatusResponse, error) {
	log.Println("GetModelStatus request")
	if _, ok := models[id.ModelId]; !ok {
		models[id.ModelId] = model{Status: pb.Status_DELETED}
	}
	return &pb.StatusResponse{
		Status: models[id.ModelId].Status,
	}, nil
}

func (s *ExampleServer) mustEmbedUnimplementedTrainToUseServer() {
	panic("implement me")
}
