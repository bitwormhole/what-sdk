package test4whatsdk
import (
    p83f08165a "github.com/bitwormhole/what-sdk/src/test/golang/units"
     "github.com/starter-go/application"
)

// type p83f08165a.Example in package:github.com/bitwormhole/what-sdk/src/test/golang/units
//
// id:com-83f08165af785d2a-units-Example
// class:
// alias:
// scope:singleton
//
type p83f08165af_units_Example struct {
}

func (inst* p83f08165af_units_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-83f08165af785d2a-units-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p83f08165af_units_Example) new() any {
    return &p83f08165a.Example{}
}

func (inst* p83f08165af_units_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p83f08165a.Example)
	nop(ie, com)

	


    return nil
}


