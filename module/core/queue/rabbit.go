package queue

import (
	"context"
	"encoding/json"
	"log"

	"cdr/module/core/model"

	"gorm.io/gorm"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	QueueName string
}

var R *Rabbit

// NewRabbit creates a Rabbit connection and declares the queue
func NewRabbit(amqpURL, queueName string) (*Rabbit, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	r := &Rabbit{conn: conn, ch: ch, QueueName: queueName}
	R = r
	return r, nil
}

// PublishCDR publishes a CDR message to the queue
func (r *Rabbit) PublishCDR(ctx context.Context, cdr model.CDR) error {
	body, err := json.Marshal(cdr)
	if err != nil {
		return err
	}
	return r.ch.PublishWithContext(ctx, "", r.QueueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

// ConsumeAndProcess starts consuming messages and upserts them into Postgres
func (r *Rabbit) ConsumeAndProcess(db *gorm.DB) error {
	msgs, err := r.ch.Consume(
		r.QueueName,
		"",
		false, // autoAck false
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var cdr model.CDR
			if err := json.Unmarshal(d.Body, &cdr); err != nil {
				log.Printf("failed to unmarshal cdr: %v", err)
				d.Nack(false, false)
				continue
			}

			// Upsert by CallID: if exists update, otherwise insert
			var existing model.CDR
			tx := db.Where("call_id = ?", cdr.CallID).Assign(cdr).FirstOrCreate(&existing)
			if tx.Error != nil {
				log.Printf("db upsert error: %v", tx.Error)
				d.Nack(false, true) // requeue on DB error
				continue
			}

			d.Ack(false)
		}
	}()
	return nil
}

// Close closes the connection and channel
func (r *Rabbit) Close() {
	if r == nil {
		return
	}
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
