package app

import (
	"fmt"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/env_utils"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func StartApplication() {
	// get the file name and line number when crash the code
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Server Option
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	RpcAPI(s)

	NETWORK := env_utils.GetEnvironmentVariable("NETWORK")
	ADDRESS := env_utils.GetEnvironmentVariable("PORT")
	lis, err := net.Listen(NETWORK, ADDRESS)

	if err != nil {
		logger_utils.Error("Error while listening Network: ", err)
	}
	go func() {
		logger_utils.Info("Starting Server...")
		if err := s.Serve(lis); err != nil {
			logger_utils.Error(fmt.Sprintln("Unable to serve : ", err), err)
		}
	}()

	// wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// BLock until signal received
	<-ch

	fmt.Println("Stop Server")
	s.Stop()
	fmt.Println("Closing the listener")
	if err := lis.Close(); err != nil {
		logger_utils.Info("Error While stopping server.")
	}
	fmt.Println("Closing program")

}
