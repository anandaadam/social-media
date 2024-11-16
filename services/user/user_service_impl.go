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

// UserServiceImpl adalah struct yang mengimplementasikan interface UserService.
type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *gorm.DB
	// Tambahkan dependensi yang diperlukan di sini, seperti repository atau logger.
}

// NewUserService membuat instance baru dari UserServiceImpl.
func NewUserService(userRepository user_repository.UserRepository, db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

// CreateUser adalah implementasi dari method CreateUser di UserService.
func (userService *UserServiceImpl) CreateUser(ctx *fiber.Ctx, userRequest *user.CreateUserRequest) (string, error) {
	// Hash password
	var email string
	password, err := helpers.HashPassword(userRequest.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	userRequest.Password = password

	// Start database transaction
	err = userService.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		email, err = userService.UserRepository.CreateUser(ctx, tx, userRequest)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	// Initialize Kafka producer
	producer, err := kafka.NewProducer(config.KafkaConfig())
	if err != nil {
		return "", fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	defer producer.Close() // Ensure producer is closed to release resources

	messageBody := map[string]string{
		"eventType": "signup_notification",
		"mediaType": "email",
		"sendTo":    email,
		"message":   fmt.Sprintf("Selamat anda berhasil signup dengan email %s", email),
	}

	var messageValue []byte
	messageValue, err = json.Marshal(messageBody)
	if err != nil {
		return "", fmt.Errorf("failed to produce Kafka message: %w", err)
	}

	// Produce Kafka message
	topic := "signup_user"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: messageValue,
	}, nil) // Delivery channel

	if err != nil {
		return "", fmt.Errorf("failed to produce Kafka message: %w", err)
	}

	return email, nil
}
