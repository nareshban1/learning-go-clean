package permission

import "go.uber.org/fx"

var Module = fx.Module("permission",
	fx.Options(
		fx.Provide(
			NewRepository,
			NewService,
			NewController,
			NewRoute,
		),

		fx.Invoke(RegisterRoute),
	))
