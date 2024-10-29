package infrastructure

import "angmorning.com/internal/services/users/domain"

func providerTypeToString(providers []domain.ProviderType) []string {
	converted := make([]string, len(providers))
	for i, p := range providers {
		converted[i] = string(p)
	}
	return converted
}
