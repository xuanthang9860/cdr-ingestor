package main

import (
	"cdr/module/core/model"
	"cdr/module/core/queue"
	"cdr/module/core/transport/http/handler"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=cdr password=admin dbname=cdr port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.CDR{})

	handler.RegisterDB(db)

	// Init RabbitMQ (assume localhost)
	amqpURL := "amqp://guest:guest@localhost:5672/"
	rb, err := queue.NewRabbit(amqpURL, "cdr_queue")
	if err != nil {
		log.Printf("warning: failed to connect rabbitmq: %v - continuing without queue", err)
	} else {
		// start consumer
		if err := rb.ConsumeAndProcess(db); err != nil {
			log.Printf("failed to start consumer: %v", err)
		}
	}

	r := gin.Default()
	handler.APICallHandler(r)

	// run server with graceful shutdown
	srvCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("server exited: %v", err)
		}
	}()

	<-srvCtx.Done()
	log.Println("shutting down")
	if rb != nil {
		rb.Close()
	}
}
