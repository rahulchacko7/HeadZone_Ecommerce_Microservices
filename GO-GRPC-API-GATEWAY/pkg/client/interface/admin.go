package interfaces

import "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
}
