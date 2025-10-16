package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}

func handleConnection(c net.Conn) {
	fmt.Print(".")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)

		randNum := random(1, 1001, r)
		resp := "Random number: " + strconv.Itoa(randNum) + "\n"
		c.Write([]byte(resp))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
