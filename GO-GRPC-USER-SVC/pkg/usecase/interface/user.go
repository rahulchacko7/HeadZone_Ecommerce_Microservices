package interfaces

import (
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/domain"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}
