package whatsdk

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/bitwormhole/what-sdk"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0

	theMainModuleResPath  = "src/main/resources"
	theTestModuleResPath  = "src/test/resources"
	theDebugModuleResPath = "src/debug/resources"
	theBenchModuleResPath = "src/benchmark/resources"
)

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/debug/resources"
var theDebugModuleResFS embed.FS

//go:embed "src/benchmark/resources"
var theBenchModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewModuleMain ...
func NewModuleMain() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(starter.Module())

	return mb
}

// NewModuleTest ...
func NewModuleTest() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}

// NewModuleDebug ...
func NewModuleDebug() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#debug")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theDebugModuleResFS, theDebugModuleResPath)
	return mb
}

// NewModuleBenchmark ...
func NewModuleBenchmark() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#benchmark")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theBenchModuleResFS, theBenchModuleResPath)
	return mb
}
