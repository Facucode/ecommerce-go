package ports

import "ecommerce-go/internal/core/domain"

type ConfigService interface {
	GetDomainEnv() domain.Environment
}
