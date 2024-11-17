package user_service

import (
	"encoding/json"
	"fmt"
	"user-service/config"
	"user-service/helpers"
	"user-service/models/user"
	user_repository "user-service/repositories/user"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository user_repository.UserRepository, db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (userService *UserServiceImpl) CreateUser(ctx *fiber.Ctx, userRequest *user.CreateUserRequest) (string, error) {
	var email string
	password, err := helpers.HashPassword(userRequest.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	userRequest.Password = password

	username := helpers.GenerateUsername(userRequest.Email)

	userInput := &user.User{}
	userInput.Email = userRequest.Email
	userInput.Password = userRequest.Password
	userInput.Username = username

	err = userService.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		email, err = userService.UserRepository.CreateUser(ctx, tx, userInput)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	_, err = userService.PublishUserEvent(email)

	if err != nil {
		return "", fmt.Errorf("publis event failed: %w", err)
	}

	return email, nil
}

func (userService *UserServiceImpl) PublishUserEvent(sendTo string) (string, error) {
	producer, err := kafka.NewProducer(config.KafkaConfig())
	if err != nil {
		return "", fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	defer producer.Close()

	messageBody := map[string]string{
		"eventType": "signup_notification",
		"mediaType": "email",
		"sendTo":    sendTo,
		"message":   fmt.Sprintf("Selamat anda berhasil signup dengan email %s", sendTo),
	}

	var messageValue []byte
	messageValue, err = json.Marshal(messageBody)
	if err != nil {
		return "", fmt.Errorf("failed to produce Kafka message: %w", err)
	}

	topic := "signup_user"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: messageValue,
	}, nil)

	if err != nil {
		return "", fmt.Errorf("failed to produce Kafka message: %w", err)
	}

	return "success", nil
}
