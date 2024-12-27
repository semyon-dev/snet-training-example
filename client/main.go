package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
	"snet-training-example/training"
	"time"
)

var HashPrefix32Bytes = []byte("\x19Ethereum Signed Message:\n32")

func getSignature(text string, blockNumber int, privateKey *ecdsa.PrivateKey) (signature []byte) {
	message := bytes.Join([][]byte{
		[]byte(text),
		crypto.PubkeyToAddress(privateKey.PublicKey).Bytes(),
		math.U256Bytes(big.NewInt(int64(blockNumber))),
	}, nil)
	hash := crypto.Keccak256(
		HashPrefix32Bytes,
		crypto.Keccak256(message),
	)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		panic(fmt.Sprintf("Cannot sign test message: %v", err))
	}
	log.Println("signed with", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	return signature
}

func main() {

	// localhost:7000 - address of daemon
	conn, err := grpc.NewClient("localhost:7000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := training.NewDaemonClient(conn) // здесь создаем именно демон клиент, а не ModelClient
	//client := training.NewModelClient(conn) // здесь создаем именно демон клиент, а не ModelClient
	//
	privateKey, err := crypto.HexToECDSA("e7638fd785fdb5cf12df0b1d7b5584cc20d4e8526403f0df105aadf23728f538")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.CreateModel(context.Background(), &training.NewModelRequest{
		Authorization: &training.AuthorizationDetails{
			CurrentBlock:  0,
			Message:       "__CreateModel",
			Signature:     getSignature("__CreateModel", 0, privateKey),
			SignerAddress: "0x9Fc6bd8e2540db7247A0772aA7eDBFA0A59d78C0",
		},
		Model: &training.NewModel{
			Name:            "test",
			Description:     "test",
			GrpcMethodName:  "stt",
			GrpcServiceName: "ProMethods",
			AddressList:     nil,
			IsPublic:        true,
			OrganizationId:  "semyon_dev",
			ServiceId:       "semyon_dev",
			GroupId:         "default_group",
		},
	})
	fmt.Println("create model resp: ", resp)
	if err != nil {
		log.Println("error", err)
		return
	}

	model, err := client.GetModel(context.Background(), &training.CommonRequest{
		Authorization: &training.AuthorizationDetails{
			CurrentBlock:  0,
			Message:       "__GetModel",
			Signature:     getSignature("__GetModel", 0, privateKey),
			SignerAddress: "0x9Fc6bd8e2540db7247A0772aA7eDBFA0A59d78C0",
		}, ModelId: resp.ModelId,
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("get model resp: ", model)

	//updateModelResp, err := client.UpdateModel(context.Background(), &training.UpdateModelRequest{
	//	Authorization: &training.AuthorizationDetails{
	//		CurrentBlock:  0,
	//		Message:       "__UpdateModel",
	//		Signature:     getSignature("__UpdateModel", 0, privateKey),
	//		SignerAddress: "0x9Fc6bd8e2540db7247A0772aA7eDBFA0A59d78C0",
	//	},
	//	ModelId:     resp.ModelId,
	//	ModelName:   "NEW NAME",
	//	Description: "UPDATED DESC",
	//	AddressList: nil,
	//})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println("update model resp: ", updateModelResp)

	//model2, err2 := client.GetModel(context.Background(), &training.CommonRequest{
	//	Authorization: &training.AuthorizationDetails{
	//		CurrentBlock:  0,
	//		Message:       "__GetModel",
	//		Signature:     getSignature("__GetModel", 0, privateKey),
	//		SignerAddress: "0x9Fc6bd8e2540db7247A0772aA7eDBFA0A59d78C0",
	//	}, ModelId: resp.ModelId,
	//})
	//if err2 != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println("[2] get model resp: ", model2)

	//time.Sleep(2 * time.Second)

	//os.Exit(0)

	price, err := client.ValidateModelPrice(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("ValidateModelPrice resp:", price)

	str, err := client.UploadAndValidate(context.Background())
	defer func(str training.Daemon_UploadAndValidateClient) {
		st, err := str.CloseAndRecv()
		if err != nil {
			log.Println(err)
		}
		if st != nil {
			log.Println(st.Status.String())
		}
	}(str)
	if err != nil {
		log.Println(err)
		return
	}

	err = str.Send(&training.UploadAndValidateRequest{
		Authorization: nil,
		UploadInput: &training.UploadInput{
			ModelId:     "1",
			Data:        []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			FileName:    "test.txt",
			FileSize:    9,
			BatchSize:   9,
			BatchNumber: 1,
			BatchCount:  1,
		},
	})
	if err != nil {
		log.Println("stream send err: ", err)
		return
	}

	time.Sleep(5 * time.Second)

	deleteModelResp, err := client.DeleteModel(context.Background(), &training.CommonRequest{
		Authorization: &training.AuthorizationDetails{
			CurrentBlock:  0,
			Message:       "__DeleteModel",
			Signature:     getSignature("__DeleteModel", 0, privateKey),
			SignerAddress: "0x9Fc6bd8e2540db7247A0772aA7eDBFA0A59d78C0",
		},
		ModelId: resp.ModelId,
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("delete model resp: %+v\n", deleteModelResp)

	conn.Close()

	//reply2, err2 := client.GetMethodMetadata(context.Background(), &training.MethodMetadataRequest{
	//	GrpcMethodName:  "stt",
	//	GrpcServiceName: "ProMethods",
	//	//ModelId: "1",
	//})
	//if err2 != nil {
	//	log.Println(err2)
	//}
	//fmt.Println(reply2)
}
