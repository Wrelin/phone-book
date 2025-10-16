package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func randomInt(min, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}

func handle(c net.Conn) {
	fmt.Print(".")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		minInt := 1
		maxInt := 1001

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

		words := strings.Fields(temp)
		if len(words) > 0 {
			i, err := strconv.Atoi(words[0])
			if err == nil {
				minInt = i
			}
		}

		if len(words) > 1 {
			i, err := strconv.Atoi(words[1])
			if err == nil {
				maxInt = i
			}
		}

		if minInt >= maxInt {
			minInt = 1
			maxInt = 1001
		}

		randNum := randomInt(minInt, maxInt, r)
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

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				if _, ok := <-interrupt; ok {
					fmt.Println(err)
				}
				return
			}
			go handle(c)
		}
	}()

	<-interrupt
	close(interrupt)

	fmt.Println("\nStop server")
}
