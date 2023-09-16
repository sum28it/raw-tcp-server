package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var (
	port int
)

// echoUpper reads from the connection and writes the capitalised text back
func echoUpper(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received: ", text)
		fmt.Fprintf(w, "%s\n", strings.ToUpper(text))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	flag.IntVar(&port, "port", 8080, "Port on which the server listens")
	flag.Parse()

	tcpListener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	if err != nil {
		log.Fatal(err)
	}
	defer tcpListener.Close()

	log.Printf("listening at localhost: %s\n", tcpListener.Addr())
	// Accept one connection at a time and process it
	for {
		conn, err := tcpListener.Accept()
		fmt.Printf("Receivded connection from, %s\n", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		go echoUpper(conn, conn)
	}

}
