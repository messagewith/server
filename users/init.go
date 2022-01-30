package users

import (
	"github.com/kamva/mgm/v3"
	database "messagewith-server/users/database"
)

var (
	Service    *service
	collection *mgm.Collection
)

func InitService() {
	collection = database.GetDB().UseCollection()
	Service = GetService(&Repository{})
}
