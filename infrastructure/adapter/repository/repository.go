package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	beer "github.com/patriciabonaldy/punkapi/domain/entity"
	ports "github.com/patriciabonaldy/punkapi/domain/ports"
	"github.com/patriciabonaldy/punkapi/infrastructure/errors"
	"github.com/streadway/amqp"
)

const (
	punkapiEndpoint = "https://api.punkapi.com/v2/beers"
)

type repository struct {
	url string
}

// NewRepository initialize csv repository
func NewRepository() ports.BeerRepo {
	return &repository{url: punkapiEndpoint}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beer.Beer, error) {
	var beers []beer.Beer
	response, err := http.Get(fmt.Sprintf("%v", r.url))

	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error obteniendo endpoint %v", r.url)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error leyendo el response %v", r.url)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error parsing to beers")
	}

	return beers, nil
}

func saveToCsv(nameFile string, records [][]string) error {
	f, err := os.Create(nameFile)
	defer f.Close()
	//CSVheader(f)
	csvWriter := csv.NewWriter(f)

	csvWriter.WriteAll(records)
	csvWriter.Flush()

	return err
}

//Rabbit struct
type Rabbit struct {
	con       beer.Conn
	queueName string
	message   []byte
}

// NewRabbitMQ initialize RabbitMQ repository
func NewRabbitMQ() ports.RabbitMq {
	return &Rabbit{queueName: ""}
}

//GetConn get conex to RabbitMQ
func (rb *Rabbit) GetConn(rabbitURL string) error {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		rb.con = beer.Conn{}
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	rb.con = beer.Conn{Channel: ch}
	return nil
}

//QueueDeclare queueName to RabbitMQ
func (rb *Rabbit) QueueDeclare(queueName string) error {
	exchangeName := "beer_Queue"

	// Create the exchange if it doesn't already exist.
	err := rb.con.Channel.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		//panic("Error creating the exchange")
	}
	// create the queue if it doesn't already exist
	//rb.queueName = queueName*/
	queue, err := rb.con.Channel.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = rb.con.Channel.QueueBind(queue.Name, "", exchangeName, false, nil)
	if err != nil {
		return err
	}
	return nil
}

//Publish into queueName to RabbitMQ
func (rb *Rabbit) Publish(routingKey string) error {
	body := amqp.Publishing{
		ContentType:  "application/json", //"application/json",
		Body:         rb.message,
		DeliveryMode: amqp.Persistent,
	}
	err := rb.con.Channel.Publish("beer_Queue", "",
		// mandatory - we don't care if there I no queue
		false,
		// immediate - we don't care if there is no consumer on the queue
		false, body)
	if err != nil {
		return err
	}
	return nil
}

//SetMessage into queueName to RabbitMQ
func (rb *Rabbit) SetMessage(message []byte) {
	rb.message = []byte(message)
}

/*//StartConsumer into queueName to RabbitMQ
func (rb Rabbit) StartConsumer(queueName, routingKey string, handler func(d amqp.Delivery) bool, concurrency int) error {

}*/
