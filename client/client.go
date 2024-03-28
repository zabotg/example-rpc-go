package client

import (
	"fmt"
	"journey-rpc/common"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

// StartClient initializes the RPC client and performs the specified operation.
func StartClient() {
	// Dial RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	// Default arguments
	firstArgument, secondArgument := 5, 5
	operation := "Sum"

	if len(os.Args) > 4 {
		operation = os.Args[2]
		firstArgument, _ = strconv.Atoi(os.Args[3])
		secondArgument, _ = strconv.Atoi(os.Args[4])
	}

	// Create arguments for RPC call
	args := &common.Args{FirstNumber: firstArgument, SecondNumber: secondArgument}
	var reply int

	// Perform the operation based on the specified operation type
	switch operation {
	case "Sum":
		reply = performOperation(client, "Calculator.Sum", args)
	case "Multiply":
		reply = performOperation(client, "Calculator.Multiply", args)
	default:
		log.Fatal("Invalid operation:", operation)
	}

	fmt.Printf("Result: %d\n", reply)
}

// performOperation calls the specified RPC method and returns the result.
func performOperation(client *rpc.Client, methodName string, args *common.Args) int {
	var reply int
	// Call RPC method
	err := client.Call(methodName, args, &reply)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	return reply
}
