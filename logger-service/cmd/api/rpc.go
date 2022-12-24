package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)


type RPCServer struct {} // the type thats going to be the rpc server


type RPCPayload struct{ /// type of data we are going to receive for any  methods that are tight to rpc server


	Name string
	Data string
}


// declare one function that has the receiver pointer to rpc server -> requiring a payload of some sort and a response ( pointer to string) making sure we can send
// a response back
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string)  error {

	collection:= client.Database("logs").Collection("logs")
	_,err:= collection.InsertOne(context.TODO(), data.LogEntry{

		Name: payload.Name,
		Data: payload.Data,
		CreatedAt:  time.Now(),
	})
	if err!=nil {
		log.Println("error writing to mongo",err)
		return err

	}


	*resp = "Processed payload via RPC:" + payload.Name

	return nil 
}

