package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/wrelin/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var min = 0
var max = 100

func random(min, max int, rnd *rand.Rand) int {
	return rnd.Intn(max-min) + min
}

// Extra function for creating secure random numbers
//
// func randomSecure(min, max int) int {
// 	v, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
// 	if err != nil {
// 		fmt.Println(err)
// 		return min
// 	}
// 	fmt.Println("**", v, min, max)
//
// 	return min + int(v.Uint64())
// }

func getString(len int64, rnd *rand.Rand) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		// For getting valid ASCII characters
		myRand := random(0, 94, rnd)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

type RandomServer struct {
	proto.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *proto.RequestDateTime) (*proto.DateTime, error) {
	currentTime := time.Now()
	response := &proto.DateTime{
		Value: currentTime.String(),
	}

	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *proto.RandomParams) (*proto.RandomInt, error) {
	rnd := rand.New(rand.NewSource(r.GetSeed()))
	place := r.GetPlace()
	temp := random(min, max, rnd)
	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(min, max, rnd)
	}

	response := &proto.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}

func (RandomServer) GetSum(ctx context.Context, r *proto.RequestSum) (*proto.ResponseSum, error) {
	response := &proto.ResponseSum{
		Sum: r.GetFirst() + r.GetSecond(),
	}

	return response, nil
}

func (RandomServer) GetRandomPass(ctx context.Context, r *proto.RequestPass) (*proto.RandomPass, error) {
	rnd := rand.New(rand.NewSource(r.GetSeed()))
	temp := getString(r.GetLength(), rnd)

	response := &proto.RandomPass{
		Password: temp,
	}

	return response, nil
}

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()
	var randomServer RandomServer
	proto.RegisterRandomServer(server, randomServer)

	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Serving requests...")
	server.Serve(listen)
}
