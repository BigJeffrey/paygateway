package middlewares

import (
	"paygateway/dao/interfaces"
)

type Middleware struct {
	Dao interfaces.AppDao
}
