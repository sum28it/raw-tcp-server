package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	port int
)

func main() {

	//Flag to specify port number
	flag.IntVar(&port, "port", 8080, "Port on which the server is listening")
	flag.Parse()

	// Connct to the server using TCP
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: port})
	if err != nil {
		log.Fatalf("error connecting to localhost: %d, %v", port, err)
	}
	defer conn.Close()

	// Read from the connection and print it on Stdout
	// in a new goroutine
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("SERVER:\t", scanner.Text())

		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading from %s, %v", conn.RemoteAddr(), err)
		}
	}()

	// Read input from Stdin
	scannerStdin := bufio.NewScanner(os.Stdin)

	for scannerStdin.Scan() {
		text := scannerStdin.Text()
		log.Printf("sent: %s\n", text)
		_, err = conn.Write([]byte(text + "\n"))
		if err != nil {
			log.Fatal(err)
		}

	}

}
