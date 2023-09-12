package service

import (
	"awesomeProject6/pkg/config/kafka"
	"awesomeProject6/pkg/dataaccess"
	"awesomeProject6/pkg/models"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
)

func CreateUser(user models.User) error {
	manager := dataaccess.MongoManager("users")
	idcount, err := manager.Totalcount(context.Background())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	user.Id = idcount + 1
	user.CreatedAt = time.Now()
	_, err = manager.Insert(context.Background(), user)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func UpdateProduct(id int) error {
	producer, err := kafka.CreateKafkaproducer()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	topic := "product-topic"
	key := "product-created"
	value := strconv.Itoa(id)
	fmt.Println(value)
	err = dataaccess.ProduceToKafka(producer, topic, key, value)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	consumer, err := kafka.CreateKafkaConsumer()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = dataaccess.Consumemessage(topic, consumer)
	if err != nil {
		return err
	}
	return nil
}
func CreateProduct(req models.Product) error {
	manager := dataaccess.MongoManager("product")
	idcount, err := manager.Totalcount(context.Background())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	req.ProductID = idcount + 1
	req.Created_At = time.Now()
	_, err = manager.Insert(context.Background(), req)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
