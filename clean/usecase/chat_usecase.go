package usecase

import (
	"errors"
	"github.com/riimi/tutorial-grpc-chat/clean/domain"
)

type ChatService interface {
	Broadcast(msg *domain.Message) error
	Recv(string) (<-chan *domain.Message, error)
	GenerateRandomToken(int) string
	RegisterUser(*domain.User) error
	UnregisterUser(*domain.User) error
	DuplicateToken(string) bool
	UserByToken(string) (*domain.User, error)
}

type ChatUsecase struct {
	service ChatService
}

func NewChatUsecase(chatService ChatService) *ChatUsecase {
	return &ChatUsecase{
		service: chatService,
	}
}

func (u *ChatUsecase) SendMessageAll(msg *domain.Message) error {
	return u.service.Broadcast(msg)
}

func (u *ChatUsecase) Subscribe(token string) (<-chan *domain.Message, error) {
	if !u.service.DuplicateToken(token) {
		return nil, errors.New("Invalid token")
	}

	return u.service.Recv(token)
}

func (u *ChatUsecase) Join(name string) (*domain.User, error) {
	var token string
	for {
		token = u.service.GenerateRandomToken(16)
		if !u.service.DuplicateToken(token) {
			break
		}
	}
	newUser := &domain.User{
		Name:  name,
		Token: token,
	}

	if err := u.service.RegisterUser(newUser); err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *ChatUsecase) Quit(token string) (*domain.User, error) {
	delUser, err := u.service.UserByToken(token)
	if err != nil {
		return nil, err
	}

	if err := u.service.UnregisterUser(delUser); err != nil {
		return delUser, err
	}

	return delUser, nil
}

func (u *ChatUsecase) ListUser() ([]*domain.User, error) {
	return []*domain.User{}, nil
}