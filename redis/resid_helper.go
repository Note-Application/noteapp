package redis

import (
	"time"
)

// SetNote stores a note in Redis with an expiry time
func SetNote(noteID string, content string) error {
	return RedisClient.Set(RedisCtx, "note:"+noteID, content, 30*time.Minute).Err()
}

// GetNote retrieves a note from Redis
func GetNote(noteID string) (string, error) {
	return RedisClient.Get(RedisCtx, "note:"+noteID).Result()
}

// DeleteNote removes a note from Redis
func DeleteNote(noteID string) error {
	return RedisClient.Del(RedisCtx, "note:"+noteID).Err()
}

// GetAllNoteKeys fetches all keys of stored notes
func GetAllNoteKeys() ([]string, error) {
	return RedisClient.Keys(RedisCtx, "note:*").Result()
}
