package service

import (
	"context"

	"github.com/roshan8/change-events/api"
	"github.com/roshan8/change-events/internal/repository"
)

type KubernetesService struct {
	// Add any dependencies here
	repo *repository.KubernetesRepository
}

func NewKubernetesService(repo *repository.KubernetesRepository) *KubernetesService {
	return &KubernetesService{repo: repo}
}

func (k *KubernetesService) CreateKubernetesEvent(ctx context.Context, kevent api.KubernetesEvent) (api.KubernetesEvent, error) {
	// Implementation
	// Add business logic validation, implement the logic to create a new Kubernetes event etc
	return k.repo.Create(ctx, &kevent)
}
