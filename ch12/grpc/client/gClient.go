package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wrelin/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = ":8080"

func AskingDateTime(ctx context.Context, m proto.RandomClient) (*proto.DateTime, error) {
	request := &proto.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}

func AskPass(ctx context.Context, m proto.RandomClient, seed int64, length int64) (*proto.RandomPass, error) {
	request := &proto.RequestPass{
		Seed:   seed,
		Length: length,
	}

	return m.GetRandomPass(ctx, request)
}

func AskRandom(ctx context.Context, m proto.RandomClient, seed int64, place int64) (*proto.RandomInt, error) {
	request := &proto.RandomParams{
		Seed:  seed,
		Place: place,
	}

	return m.GetRandom(ctx, request)
}

func AskSum(ctx context.Context, m proto.RandomClient, first, second int64) (*proto.ResponseSum, error) {
	request := &proto.RequestSum{
		First:  first,
		Second: second,
	}

	return m.GetSum(ctx, request)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	seed := int64(rand.Intn(100))

	client := proto.NewRandomClient(conn)
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Date and Time:", r.Value)

	length := int64(rnd.Intn(20))
	p, err := AskPass(context.Background(), client, 100, length+1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Password:", p.Password)

	place := int64(rnd.Intn(100))
	i, err := AskRandom(context.Background(), client, seed, place)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 1:", i.Value)

	k, err := AskRandom(context.Background(), client, seed, place-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 2:", k.Value)

	sum, err := AskSum(context.Background(), client, i.Value, k.Value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum:", sum.Sum)
}
