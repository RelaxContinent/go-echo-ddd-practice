package initial

import (
	"github.com/RelaxContinent/go-echo-ddd-practice/initial/application"
	"github.com/RelaxContinent/go-echo-ddd-practice/initial/infrastructure"
	"github.com/RelaxContinent/go-echo-ddd-practice/initial/presentation"
	"go.uber.org/fx"
)

var Modules = fx.Module(
	"InitialModules",
	fx.Provide(
		infrastructure.NewInitialRepositoryImpl,
		application.NewInitialUseCase,
		presentation.NewInitialHandler,
	),
	fx.Invoke(
		presentation.SetUpRoute,
	),
)
