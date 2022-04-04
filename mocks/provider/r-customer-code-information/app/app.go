package app

import (
	"log"

	"github.com/danilotadeu/services/r-customer-code-information-provider/app/codeinformation"
)

//Container ...
type Container struct {
	CodeInformation codeinformation.App
}

//Register app container
func Register() *Container {
	container := &Container{
		// CodeInformation: codeinformation.NewApp(),
	}
	log.Println("Registered -> App")
	return container
}
