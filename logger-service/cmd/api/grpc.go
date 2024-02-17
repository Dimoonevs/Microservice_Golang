package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/Microservice_Golang/logger/data"
	"github.com/Dimoonevs/Microservice_Golang/logger/logs"
	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogSerivceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// write response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	logs.RegisterLogSerivceServer(s, &LogServer{Models: app.Models})
	log.Printf("grps server started on port %s", gRpcPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
