package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func sendTCPRequest(serverIP string, serverPort int) { 
	address := fmt.Sprintf("%s:%d", serverIP, serverPort)
	for {
		start := time.Now()
		conn, err := net.Dial("tcp", address)
		if err != nil {
			color.Red("Connection timed out") 
			time.Sleep(1 * time.Second) 
			continue 
		}
		defer conn.Close()
		message := []byte("GET / HTTP/1.1\r\nHost: " + serverIP + "\r\n\r\n")
		_, err = conn.Write(message)
		if err != nil {
			color.Red("Erreur d'envoi du message") 
		} else {
			duration := time.Since(start)
			green := color.New(color.FgGreen) 
			fmt.Printf("Ping to %s : %s ms port = %s\n", green.Sprint(serverIP), green.Sprint(int(duration.Milliseconds())), green.Sprint(serverPort))
		}
		time.Sleep(1 * time.Second) 
	}
}

func main() {
	serverIP := flag.String("host", "0.0.0.0", "Adresse IP")
	serverPort := flag.Int("port", 80, "Port à utiliser")
	flag.Parse()
	sendTCPRequest(*serverIP, *serverPort)
}
