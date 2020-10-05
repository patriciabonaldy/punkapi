package ports

import (
	beer "github.com/patriciabonaldy/punkapi/domain/entity"
)

// Fetching provides beer fetching operations
type Fetching interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beer.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beer.Beer, error)
}

//RabbitMq provides operations RabbitMq
type RabbitMq interface {
	GetConn(rabbitURL string) error
	QueueDeclare(queueName string) error
	SetMessage(message []byte)
	Publish(routingKey string) error
	//StartConsumer(queueName, routingKey string, handler func(d amqp.Delivery) bool, concurrency int) error
}
