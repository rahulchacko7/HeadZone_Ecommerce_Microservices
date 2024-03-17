package interfaces

import "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignUp) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
