package main

import (
	"bufio"
	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, err := os.Open("../wordlist")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Read and process each word
	for scanner.Scan() {
		word := scanner.Text()
		request := stubs.Request{Message: word}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responsed:" + response.Message)
	}

	//TODO: connect to the RPC server and send the request(s)

}
