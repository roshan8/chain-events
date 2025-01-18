package repository

import (
	"context"

	"github.com/roshan8/chain-events/api"
	"gorm.io/gorm"
)

type KubernetesRepository struct {
	db *gorm.DB
}

func NewKubernetesRepository(db *gorm.DB) *KubernetesRepository {
	return &KubernetesRepository{db: db}
}

func (k *KubernetesRepository) Create(ctx context.Context, kevent *api.KubernetesEvent) (api.KubernetesEvent, error) {
	// Implementation
	return *kevent, nil
}
