package service

import (
	"E-Commerce-Golang/internal/domain"
	"E-Commerce-Golang/internal/dto"
	"E-Commerce-Golang/internal/repository"
	"errors"
	"fmt"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) FindUserByemail(email string) (*domain.User, error) {

	//Perform db operation and business logic
	return nil, nil
}

func (s UserService) Signup(input dto.UserSignuP) (string, error) {
	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// Generate toekn
	userInfo := fmt.Sprintf("%v %v %v", user.ID, user.Email, user.UserType)
	return userInfo, err
}
func (s UserService) findUserByemail(email string) (*domain.User, error) {
	// business logic
	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByemail(email)
	// compare password and generate token
	if err != nil {
		return "", errors.New("user does not exists with the the provided emial id")
	}
	return user.Email, err
}
func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}
func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}
func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "nil", nil
}
func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) CreateOrder(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}
func GetOrderById(id uint, uid int) (interface{}, error) {
	return nil, nil
}
