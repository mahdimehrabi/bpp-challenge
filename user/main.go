package main

import (
	"fmt"
	"go.uber.org/fx"
	"user/app/infrastractures"
)

var BootstrapModule = fx.Options(
	infrastractures.Module,
	fx.Invoke(Bootstrap),
)

func Bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastractures.PasargadLogger,
) {

}

func main() {
	fmt.Println("Hey its user service")
}
