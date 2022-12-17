package main

import (
	"github/kameshsampath/devfest-ahm22/pkg/greeter"
	"github/kameshsampath/devfest-ahm22/pkg/greeter/impl"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &impl.LinguaGreeterServer{})
	reflection.Register(s)
	log.Printf("Server started at %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
