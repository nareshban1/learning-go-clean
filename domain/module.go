package domain

import (
	"clean-architecture/domain/permission"
	"clean-architecture/domain/role"
	"clean-architecture/domain/user"

	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Options(
		user.Module,
		role.Module,
		permission.Module,
	),
)
