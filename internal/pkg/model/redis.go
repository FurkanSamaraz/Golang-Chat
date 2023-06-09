package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	api_structure "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/structures"

	"github.com/go-redis/redis/v8"
)

type RedisService struct{ Client *redis.Client }

func (redisC *RedisService) RegisterNewUser(user *api_structure.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("json", err)
		return err
	}

	err = redisC.Client.Set(context.Background(), user.Username, data, 0).Err()
	if err != nil {
		log.Println("error while adding new user", err)
		return err
	}

	return nil
}

func (redisC *RedisService) IsUserExist(username string) (bool, error) {
	exists, err := redisC.Client.Exists(context.Background(), username).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (redisC *RedisService) IsUserAuthentic(user *api_structure.User) error {
	data, err := redisC.Client.Get(context.Background(), user.Username).Bytes()
	if err != nil {
		log.Fatal(err)
	}

	var savedUser api_structure.User
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		log.Fatal(err)
	}

	if savedUser.Password != user.Password {
		log.Fatal("Invalid credentials")
	}

	return nil
}

func (redisC *RedisService) FetchChatBetween(username1, username2 string) ([]api_structure.Chat, error) {
	// Sohbet mesajları için Redis anahtarını oluşturun
	chatKey := fmt.Sprintf("chats:%s:%s", username1, username2)
	var chatHistory []api_structure.Chat

	// Belirtilen zaman aralığında sohbet mesajlarını alın
	chatData, err := redisC.Client.Get(context.Background(), chatKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Chat history not found")
		} else {
			fmt.Println("Failed to fetch chat history:", err)
		}
	}
	err = json.Unmarshal(chatData, &chatHistory)
	if err != nil {
		log.Fatal(err)
	}
	return chatHistory, nil
}

// Kullanıcının Kişi Listesini Getir. Kişiye gönderilen ve kişi tarafından alınan tüm mesajları içerir.
// Bir kişiyle son aktiviteye göre sıralanmış bir liste döndürür
func (redis *RedisService) FetchContactList(username string) ([]string, error) {
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Kişi listesini Redis'ten al
	contactList, err := redis.Client.SMembers(context.Background(), contactListKey).Result()
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

func (redis *RedisService) AddToContactList(username, contactUsername string) error {
	// Kişi listesi için Redis anahtarını oluşturun
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Kişiyi kişi listesine ekle
	err := redis.Client.SAdd(context.Background(), contactListKey, contactUsername).Err()
	if err != nil {
		return err
	}

	return nil
}

// Sohbeti kaydet
func (redisC *RedisService) SaveChatHistory(msg api_structure.Chat) error {
	chatKey := fmt.Sprintf("chats:%s:%s", msg.From, msg.To)

	chatHistory, err := redisC.Client.Get(context.Background(), chatKey).Bytes()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("error fetching chat history: %v", err)
	}

	var messages []api_structure.Chat
	if len(chatHistory) > 0 {
		err = json.Unmarshal(chatHistory, &messages)
		if err != nil {
			return fmt.Errorf("error unmarshaling chat history: %v", err)
		}
	}

	messages = append(messages, msg)

	data, err := json.Marshal(messages)
	if err != nil {
		return fmt.Errorf("error marshaling chat history: %v", err)
	}

	err = redisC.Client.Set(context.TODO(), chatKey, data, 0).Err()
	if err != nil {
		return fmt.Errorf("error saving chat history: %v", err)
	}

	return nil
}
