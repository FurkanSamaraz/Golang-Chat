package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	setup "main/internal/pkg/config"
	api_structure "main/internal/pkg/structures"

	"github.com/go-redis/redis/v8"
)

func RegisterNewUser(user *api_structure.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("json", err)
		return err
	}

	err = setup.RedisClient.Set(context.Background(), user.Username, data, 0).Err()
	if err != nil {
		log.Println("error while adding new user", err)
		return err
	}

	return nil
}

func IsUserExist(username string) (bool, error) {
	exists, err := setup.RedisClient.Exists(context.Background(), username).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func IsUserAuthentic(user *api_structure.User) error {
	data, err := setup.RedisClient.Get(context.Background(), user.Username).Bytes()
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

func FetchChatBetween(username1, username2 string) ([]api_structure.Chat, error) {
	// Sohbet mesajları için Redis anahtarını oluşturun
	chatKey := fmt.Sprintf("chats:%s:%s", username1, username2)
	var chatHistory []api_structure.Chat

	// Belirtilen zaman aralığında sohbet mesajlarını alın
	chatData, err := setup.RedisClient.Get(context.Background(), chatKey).Bytes()
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
func FetchContactList(username string) ([]string, error) {
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Kişi listesini Redis'ten al
	contactList, err := setup.RedisClient.SMembers(context.Background(), contactListKey).Result()
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

func AddToContactList(username, contactUsername string) error {
	// Kişi listesi için Redis anahtarını oluşturun
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Kişiyi kişi listesine ekle
	err := setup.RedisClient.SAdd(context.Background(), contactListKey, contactUsername).Err()
	if err != nil {
		return err
	}

	return nil
}

// Sohbeti kaydet
func SaveChatHistory(msg api_structure.Chat) error {
	chatKey := fmt.Sprintf("chats:%s:%s", msg.From, msg.To)

	chatHistory, err := setup.RedisClient.Get(context.Background(), chatKey).Bytes()
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

	err = setup.RedisClient.Set(context.TODO(), chatKey, data, 0).Err()
	if err != nil {
		return fmt.Errorf("error saving chat history: %v", err)
	}

	return nil
}
