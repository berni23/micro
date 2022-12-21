package event


import amqp "github.com/rabbitmq/amqp091-go"




func declareExchange (channel *amqp.Channel ) error {


	return channel.ExchangeDeclare(

		"logs_topic",   //name
		"topic",        //type
		true,           //durable?
		false ,        //auto-deleted? 
		false,         // internal?
		false,         // internal?
		nil  ,          //args?
	)
}


func declareRandomQueue(ch * amqp.Channel) (amqp.Queue, error) {

	return ch.QueueDeclare(
		"", //name? 
		false, // durable?
		false , // delete when unused?
		true , // exclusive
		false, // no wait ? 
		nil , // args? 
	)
}