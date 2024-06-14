package main4whatsdk
import (
    p095252e47 "github.com/bitwormhole/what-sdk/v1/implements"
     "github.com/starter-go/application"
)

// type p095252e47.Example in package:github.com/bitwormhole/what-sdk/v1/implements
//
// id:com-095252e4725e9ab3-implements-Example
// class:
// alias:
// scope:singleton
//
type p095252e472_implements_Example struct {
}

func (inst* p095252e472_implements_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-095252e4725e9ab3-implements-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p095252e472_implements_Example) new() any {
    return &p095252e47.Example{}
}

func (inst* p095252e472_implements_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p095252e47.Example)
	nop(ie, com)

	


    return nil
}


