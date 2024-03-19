package interfaces

import (
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/domain"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/utils/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignUp) (models.UserDetails, error)
	FindUserByEmail(user models.UserLogin) (models.UserDetail, error)
}
