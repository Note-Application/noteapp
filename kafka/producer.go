package kafka

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"noteapp/internal/models"
	"noteapp/redis"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer struct to manage Kafka operations
type KafkaProducer struct {
	Writer *kafka.Writer
}

// NewProducer initializes a Kafka producer
func NewProducer() *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "update-notes-topic",
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{Writer: writer}
}

// Close shuts down the Kafka writer
func (kp *KafkaProducer) Close() {
	if kp.Writer != nil {
		log.Println("🔌 Closing Kafka producer...")
		kp.Writer.Close()
	}
}

// ProduceMessage sends a message to a Kafka topic
func (kp *KafkaProducer) ProduceMessage(topic, message string) error {
	msg := kafka.Message{
		Key:   []byte(strconv.FormatInt(time.Now().UnixNano(), 10)), // Unique key per message
		Value: []byte(message),
	}

	err := kp.Writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("❌ Failed to send Kafka message: %v", err)
		return err
	}

	log.Println("✅ Message sent to Kafka:", message)
	return nil
}

// PushUpdatesToKafka fetches batched updates from Redis and pushes them to Kafka
func (kp *KafkaProducer) PushUpdatesToKafka(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Stopping Redis-to-Kafka updates...")
			return
		case <-ticker.C:
			// Get all note keys from Redis
			keys, err := redis.RedisClient.Keys(ctx, "note:*").Result()
			if err != nil {
				log.Printf("❌ Error fetching keys from Redis: %v", err)
				continue
			}

			if len(keys) == 0 {
				continue // No updates, skip this cycle
			}

			batch := make(map[string]string) // Store latest updates only

			// Collect latest updates from Redis
			for _, key := range keys {
				noteID := strings.TrimPrefix(key, "note:")
				data, err := redis.RedisClient.Get(ctx, key).Result()
				if err != nil {
					log.Printf("❌ Failed to get note %s from Redis: %v", noteID, err)
					continue
				}
				batch[noteID] = data // Keep only the latest value
			}

			// Send batch updates to Kafka
			for noteID, data := range batch {
				noteid, err := strconv.Atoi(noteID)
				if err != nil {
					log.Printf("❌ Failed to parse note ID: %v", err)
					continue
				}

				var noteData map[string]string
				err = json.Unmarshal([]byte(data), &noteData)
				if err != nil {
					log.Printf("❌ Failed to deserialize note data: %v", err)
					continue
				}

				// Convert request to note model
				note := models.Note{
					ID:      noteid,
					Title:   noteData["title"],
					Content: noteData["content"],
				}

				// Convert note to JSON for Kafka
				message, err := json.Marshal(note)
				if err != nil {
					continue
				}

				err = kp.ProduceMessage("update-notes-topic", string(message))
				if err != nil {
					log.Printf("❌ Failed to send update to Kafka: %v", err)
				} else {
					// Remove successfully processed keys from Redis
					redis.RedisClient.Del(ctx, "note:"+noteID)
				}
			}

			log.Printf("📤 Pushed %d updates to Kafka", len(batch))
		}
	}
}
