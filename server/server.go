package server

import (
	"log"
	"net"
	"net/rpc"

	"journey-rpc/common"
)

// Calculator is a type that implements calculator operations.
type Calculator int

// Multiply performs the multiplication of two numbers provided in the arguments.
func (c *Calculator) Multiply(args *common.Args, reply *int) error {
	log.Printf("Multiply called with args: %+v", args)
	*reply = args.FirstNumber * args.SecondNumber
	return nil
}

// Sum performs the addition of two numbers provided in the arguments.
func (c *Calculator) Sum(args *common.Args, reply *int) error {
	log.Printf("Sum called with args: %+v", args)
	*reply = args.FirstNumber + args.SecondNumber
	return nil
}

// StartServer initializes and starts the RPC server.
func StartServer() {
	calc := new(Calculator)

	// Register the Calculator instance for RPC.
	err := rpc.Register(calc)
	if err != nil {
		log.Fatalf("Error registering RPC service: %v", err)
	}

	const port = ":1234"

	// Listen for TCP connections on the specified port.
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	log.Printf("Server RPC listening on port %s", port)

	// Accept incoming RPC requests.
	rpc.Accept(listener)

}
