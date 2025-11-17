package framework

import (
	"go.uber.org/fx"
)

var Modules = fx.Module(
	"FrameworkModules",
	fx.Provide(
		NewLogger,
		NewCustomBinder,
		NewCustomValidator,
		NewEcho,
	),
)
