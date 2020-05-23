//+build wireinject

package main

import (
	"billme/api/inventory"

	"github.com/google/wire"
)

func InitInventoryAPI() inventory.API {
	wire.Build(
		inventory.ProvideInventoryRepo,
		inventory.ProvideInventoryService,
		inventory.ProvideInventoryAPI,
	)

	return inventory.API{}
}
