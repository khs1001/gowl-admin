package facades

import (
	"log"

	admin "github.com/khs1001/gowl-admin"
	"github.com/khs1001/gowl-admin/contracts"
)

func Admin() contracts.Admin {
	instance, err := admin.App.Make(admin.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Admin)
}
