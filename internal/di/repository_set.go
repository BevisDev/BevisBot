package di

import (
	"github.com/BevisDev/BevisBot/internal/repository"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repository.NewTaskRepository,
)
