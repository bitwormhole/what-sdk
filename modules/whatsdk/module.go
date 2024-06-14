package whatsdk

import (
	whatsdk "github.com/bitwormhole/what-sdk"
	"github.com/bitwormhole/what-sdk/gen/benchmark4whatsdk"
	"github.com/bitwormhole/what-sdk/gen/debug4whatsdk"
	"github.com/bitwormhole/what-sdk/gen/main4whatsdk"
	"github.com/bitwormhole/what-sdk/gen/test4whatsdk"
	"github.com/starter-go/application"
)

// ModuleMain ...
func ModuleMain() application.Module {

	mb := whatsdk.NewModuleMain()
	mb.Components(main4whatsdk.ExportComponents)
	return mb.Create()
}

// ModuleTest ...
func ModuleTest() application.Module {

	mb := whatsdk.NewModuleMain()
	mb.Components(test4whatsdk.ExportComponents)
	mb.Depend(ModuleMain())
	return mb.Create()
}

// ModuleDebug ...
func ModuleDebug() application.Module {

	mb := whatsdk.NewModuleDebug()
	mb.Components(debug4whatsdk.ExportComponents)
	mb.Depend(ModuleMain())
	return mb.Create()
}

// ModuleBenchmark ...
func ModuleBenchmark() application.Module {

	mb := whatsdk.NewModuleBenchmark()
	mb.Components(benchmark4whatsdk.ExportComponents)
	mb.Depend(ModuleMain())
	return mb.Create()
}
