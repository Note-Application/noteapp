package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"noteapp/internal/models"
	"noteapp/internal/services"
	"noteapp/proto/generated"
	"noteapp/redis"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteHandler struct {
	generated.UnimplementedNoteServiceServer
}

// CreateNote handles creating a new note
func (s *NoteHandler) CreateNote(ctx context.Context, req *generated.CreateNoteRequest) (*generated.NoteResponse, error) {
	note := models.Note{
		UserID:  int(req.UserId),
		Title:   req.Title,
		Content: req.Content,
	}

	// Create the note and get the ID
	id, err := services.CreateNote(note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create note")
	}

	// Set the ID in the note object
	note.ID = id

	// Return the created note
	return &generated.NoteResponse{
		Note: &generated.Note{
			Id:      int32(id),
			UserId:  req.UserId,
			Title:   req.Title,
			Content: req.Content,
		},
	}, nil
}

// GetNoteByID retrieves a note by its ID
func (s *NoteHandler) GetNoteByID(ctx context.Context, req *generated.GetNoteByIDRequest) (*generated.NoteResponse, error) {
	note, err := services.GetNoteByID(fmt.Sprintf("%d", req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Note not found")
	}

	return &generated.NoteResponse{
		Note: &generated.Note{
			Id:      int32(note.ID),
			UserId:  int32(note.UserID),
			Title:   note.Title,
			Content: note.Content,
		},
	}, nil
}

// GetNotesByUserID retrieves all notes for a specific user
func (s *NoteHandler) GetNotesByUserID(ctx context.Context, req *generated.GetNotesByUserIDRequest) (*generated.GetNotesByUserIDResponse, error) {
	notes, err := services.GetNotesByUserID(fmt.Sprintf("%d", req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "No notes found")
	}

	// Convert notes to gRPC format
	var grpcNotes []*generated.Note
	for _, note := range notes {
		grpcNotes = append(grpcNotes, &generated.Note{
			Id:      int32(note.ID),
			UserId:  int32(note.UserID),
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return &generated.GetNotesByUserIDResponse{Notes: grpcNotes}, nil
}

// UpdateNote publishes the update request to Kafka instead of updating directly
func (s *NoteHandler) UpdateNote(ctx context.Context, req *generated.UpdateNoteRequest) (*generated.NoteResponse, error) {

	// Create a structured JSON object
	noteData := map[string]string{
		"title":   req.Title,
		"content": req.Content,
	}

	// Convert the map to JSON format
	noteJSON, err := json.Marshal(noteData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to serialize note data")
	}

	// Store the updated note in Redis (with 30 seconds expiration)
	redisKey := fmt.Sprintf("%d", req.Id)
	redisValue := string(noteJSON)
	err = redis.SetNote(redisKey, redisValue)
	if err != nil {
		log.Println("Failed to store note update in Redis:", err)
		return nil, status.Errorf(codes.Internal, "Failed to store update in Redis")
	}

	log.Println("Stored note update in Redis for", req.Id)

	// Return success response immediately (actual update happens via Kafka consumer)
	return &generated.NoteResponse{
		Note: &generated.Note{
			Id:      req.Id,
			Title:   req.Title,
			Content: req.Content,
		},
	}, nil
}

// DeleteNote deletes a note by its ID
func (s *NoteHandler) DeleteNote(ctx context.Context, req *generated.DeleteNoteRequest) (*generated.DeleteNoteResponse, error) {
	err := services.DeleteNote(fmt.Sprintf("%d", req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete note")
	}

	return &generated.DeleteNoteResponse{
		Message: "Note deleted successfully",
	}, nil
}
