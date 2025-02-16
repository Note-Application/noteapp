package handlers

import (
	"context"
	"noteapp/internal/models"
	"noteapp/internal/services"
	"noteapp/proto/generated"
)

type UserHandler struct {
	generated.UnimplementedUserServiceServer
}

// CreateUser handles creating a new user
func (s *UserHandler) CreateUser(ctx context.Context, req *generated.CreateUserRequest) (*generated.UserResponse, error) {
	user := models.User{
		Email:      req.Email,
		Name:       req.Name,
		ProfilePic: req.ProfilePic,
	}

	// Create the user and get the ID
	id, err := services.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Set the ID in the user object
	user.ID = id

	// Return the response with a properly formatted User message
	return &generated.UserResponse{
		User: &generated.User{
			Id:         int32(id),
			Email:      req.Email,
			Name:       req.Name,
			ProfilePic: req.ProfilePic,
		},
	}, nil
}

// GetUserByEmail retrieves a user by their email
func (s *UserHandler) GetUserByEmail(ctx context.Context, req *generated.GetUserByEmailRequest) (*generated.UserResponse, error) {
	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return &generated.UserResponse{
		User: &generated.User{
			Id:         int32(user.ID),
			Email:      user.Email,
			Name:       user.Name,
			ProfilePic: user.ProfilePic,
		},
	}, nil
}

// UpdateUser updates an existing user
func (s *UserHandler) UpdateUser(ctx context.Context, req *generated.UpdateUserRequest) (*generated.UserResponse, error) {
	user := models.User{
		Email:      req.Email,
		Name:       req.Name,
		ProfilePic: req.ProfilePic,
	}

	err := services.UpdateUser(req.Email, user)
	if err != nil {
		return nil, err
	}

	return &generated.UserResponse{
		User: &generated.User{
			Id:         int32(user.ID),
			Email:      user.Email,
			Name:       user.Name,
			ProfilePic: user.ProfilePic,
		},
	}, nil
}

// DeleteUser deletes a user by their email
func (s *UserHandler) DeleteUser(ctx context.Context, req *generated.DeleteUserRequest) (*generated.DeleteUserResponse, error) {
	err := services.DeleteUser(req.Email)
	if err != nil {
		return nil, err
	}

	return &generated.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
