package main

import (
	"cdr/internal/config"
	"cdr/internal/model"
	"cdr/internal/queue"
	"cdr/internal/transport/http/handler"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cc-integration-team/cc-pkg/v2/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
	logger.SetDefaultLogger(logger.NewZerologAdapter(cfg.Logger))

	postgresCfg := cfg.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		postgresCfg.Host, postgresCfg.User, postgresCfg.Password, postgresCfg.DBName, postgresCfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&model.CDR{})
	handler.RegisterDB(db)

	// Init RabbitMQ (assume localhost)

	rabbitCfg := cfg.RabbitMQ
	amqpURL := rabbitCfg.URL
	rb, err := queue.NewRabbit(amqpURL, rabbitCfg.QueueName)
	if err != nil {
		logger.Fatalf("failed to connect rabbitmq: %v", err)
	} else {
		// start consumer
		if err := rb.ConsumeAndProcess(db); err != nil {
			logger.Fatalf("failed to start consumer: %v", err)
		}
	}

	r := gin.Default()
	handler.APICallHandler(r)

	// run server with graceful shutdown
	srvCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := r.Run(":8080"); err != nil {
			logger.Fatalf("failed to start server: %v", err)
		}
	}()

	<-srvCtx.Done()
	logger.Infof("shutting down")
	if rb != nil {
		rb.Close()
	}
}
