package service

import (
	"github.com/KawashiroNitori/MoeManager/internal/manager"
	"github.com/kardianos/service"
	"github.com/samber/lo"
)

var svcConfig = &service.Config{
	Name:        "MoeManager",
	DisplayName: "MoeManager",
	Description: "MoeManager is a picture manager.",
}

func NewService() service.Service {
	options := make(service.KeyValue)
	options["DelayedAutoStart"] = true
	cfg := &service.Config{
		Name:        "MoeManager",
		DisplayName: "MoeManager",
		Description: "MoeManager is a picture manager.",
		Option:      options,
	}
	m := manager.NewManager()
	return lo.Must(service.New(m, cfg))
}
