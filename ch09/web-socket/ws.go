package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"

func random(min, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket!")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("From", r.Host, "read", err)
			break
		}

		text := string(message)
		log.Print("Received: ", text)

		words := strings.Fields(text)
		if len(words) > 0 {
			n, err := strconv.Atoi(words[0])
			if err == nil {
				for i := 0; i < n; i++ {
					randNum := random(1, 1001, rnd)
					err = ws.WriteMessage(mt, []byte("Rand "+strconv.Itoa(i)+": "+strconv.Itoa(randNum)))
					if err != nil {
						log.Println("WriteMessage:", err)
						return
					}
				}
			}
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("WriteMessage:", err)
			break
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/ws", http.HandlerFunc(wsHandler))

	log.Println("Listening to TCP Port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
