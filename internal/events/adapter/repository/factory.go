package repository

import (
	"events-api/internal/events"
	"events-api/pkg/server"
)

func NewRepository(client server.RestClient) events.Repository {
	return NewRestRepository(client)
}
