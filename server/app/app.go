package app

import (
	"github.com/04Akaps/kafka-go/config"
	"github.com/04Akaps/kafka-go/server/network"
	"github.com/04Akaps/kafka-go/server/repository"
	"github.com/04Akaps/kafka-go/server/service"
)

type App struct {
	config *config.Config

	repository *repository.Repository
	service    service.ServiceImpl
	network    *network.Network
}

func NewApp(config *config.Config) error {
	a := &App{config: config}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else {
		a.service = service.NewService(config, a.repository)
		a.network = network.NewNetwork(config, a.service)

		return a.network.StartServer()
	}

}
