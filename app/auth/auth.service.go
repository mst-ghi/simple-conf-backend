package auth

import (
	"sync"
	"video-conf/core"
	"video-conf/database/models"
	"video-conf/database/repositories"
)

type AuthServiceInterface interface {
	Login(dto LoginDto) (Tokens, core.Error)
	Register(dto RegisterDto) core.Error
	Refresh(dto RefreshDto) (Tokens, core.Error)
	Logout(token string)
}

type AuthService struct {
	userRepository  repositories.UserRepositoryInterface
	tokenRepository repositories.TokenRepositoryInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository:  repositories.NewUserRepository(),
		tokenRepository: repositories.NewTokenRepository(),
	}
}

func (self *AuthService) Login(dto LoginDto) (Tokens, core.Error) {
	var tokens Tokens

	user := self.userRepository.FindByEmail(dto.Email)

	if user.ID == "" || !user.CheckPasswordHash(dto.Password) {
		return tokens, core.Error{"email": "Invalid input data"}
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go self.tokenRepository.Clearing(user.ID, &waitGroup)

	accessToken, refreshToken, expiresAt := self.tokenRepository.Create(user.ID)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.ExpiresAt = expiresAt

	waitGroup.Wait()
	return tokens, nil
}

func (self *AuthService) Register(dto RegisterDto) core.Error {
	user := self.userRepository.FindByEmail(dto.Email)

	if user.ID != "" {
		return core.Error{"email": "User exist with this email"}
	}

	user = models.User{Name: dto.Name, Email: dto.Email, Password: dto.Password}
	self.userRepository.Create(user)

	return nil
}

func (self *AuthService) Refresh(dto RefreshDto) (Tokens, core.Error) {
	var tokens Tokens

	token := self.tokenRepository.FindByRefreshAndAccess(dto.AccessToken, dto.RefreshToken)

	if token.ID == "" {
		return tokens, core.Error{"access_token": "Access token is invalid", "refresh_token": "Refresh token is invalid"}
	}

	user := token.User

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go self.tokenRepository.Clearing(user.ID, &waitGroup)

	accessToken, refreshToken, expiresAt := self.tokenRepository.Create(user.ID)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.ExpiresAt = expiresAt

	waitGroup.Wait()

	self.tokenRepository.DeleteByAccess(token.AccessToken)
	return tokens, nil
}

func (self *AuthService) Logout(token string) {
	self.tokenRepository.DeleteByAccess(token)
}
