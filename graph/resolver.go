package graph

import "github.com/emejotaw/users-api-bff/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService    *service.UserService
	addressService *service.AddressService
}

func NewResolver(userService *service.UserService, addressService *service.AddressService) *Resolver {
	return &Resolver{
		userService:    userService,
		addressService: addressService,
	}
}
