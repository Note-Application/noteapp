package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"noteapp/internal/models"
	"noteapp/internal/services"

	"github.com/segmentio/kafka-go"
)

// StartConsumer listens for note update messages
func StartConsumer(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "update-notes-topic",
		GroupID: "note-update-group",
	})

	defer func() {
		log.Println("🔌 Closing Kafka consumer...")
		reader.Close()
	}()

	log.Println("🚀 Kafka Consumer started...")

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Stopping Kafka Consumer...")
			return
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Println("❌ Error reading Kafka message:", err)
				continue
			}

			var note models.Note
			err = json.Unmarshal(msg.Value, &note)
			if err != nil {
				log.Println("❌ Failed to unmarshal note:", err)
				continue
			}

			// Update note in database
			err = services.UpdateNote(fmt.Sprintf("%d", note.ID), note)
			if err != nil {
				log.Println("❌ Failed to update note in DB:", err)
			} else {
				log.Println("✅ Note updated successfully:", note.ID)
			}
		}
	}
}
