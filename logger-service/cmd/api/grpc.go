package main

import (
	"context"
	"log-service/data"
	"log-service/logs"
)

// server -> going to receive requests

type LogServer struct {
	logs.UnimplementedLogServiceServer // required for pretty much every service written in grpc ( in order to ensure backwards compatibility)

	Models data.Models // necessary methods to wrtie to mongo

}

//
func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {

	input := req.GetLogEntry() // get log entry, it gets my input  ( input.Name, input.Data )=> specified already in logs proto

	// write the log

	logEntry := data.LogEntry{

		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)

	if err != nil {

		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return response

	res := &logs.LogResponse{Result: "logged!"}

	return res, nil
}
