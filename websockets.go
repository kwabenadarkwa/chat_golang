package main

// figure out how to convert this into a chat app that you can test with postman

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
		}
		for {
			// read the message from the browser
			// fmt.Println("this is the body of the request",r.Body)
			msgType, msg, err := conn.ReadMessage()
			// fmt.Println("message type",msgType)
			if err != nil {
				fmt.Println(err)
			}
			// print the message received to the console
			// fmt.Println("type of remote address",reflect.TypeOf(conn.RemoteAddr()))
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// write the messsage back to the browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}

		}
	})
	// for messaging you add a message as well as the intended ip address
	// then the server receives the message that you have sent and then forwards it to the intended receiver

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
		}
		for {
			// get the message
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
			}
			// decode the ip out of it to forward the message to
			var targetIP, message string
			if i := strings.Index(string(msg), ","); i >= 0 {
				_, message = string(msg)[:i], string(msg)[i:]
			} else {
				fmt.Println("wrong format")
			}
			fmt.Println("parsing done", targetIP, message)
			// close connection with other ip address
			// conn.Close()
			// open connection with new ip address and send message
			// forward the message to that ip
			tcpAddr, err := net.ResolveTCPAddr("tcp", targetIP)
			if err != nil {
				fmt.Println(err)
			}

			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				fmt.Println(err)
			}
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
			}
			var buf [512]byte
			//
			_, err = conn.Read(buf[0:])
			fmt.Printf("this is things %s", buf)
			if err != nil {
				fmt.Println(err)
			}

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
