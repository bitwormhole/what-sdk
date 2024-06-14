package debug4whatsdk
import (
    p7e252d7cb "github.com/bitwormhole/what-sdk/src/debug/golang/units"
     "github.com/starter-go/application"
)

// type p7e252d7cb.TryConnectAsSimple in package:github.com/bitwormhole/what-sdk/src/debug/golang/units
//
// id:com-7e252d7cb97f6e6c-units-TryConnectAsSimple
// class:
// alias:
// scope:singleton
//
type p7e252d7cb9_units_TryConnectAsSimple struct {
}

func (inst* p7e252d7cb9_units_TryConnectAsSimple) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7e252d7cb97f6e6c-units-TryConnectAsSimple"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7e252d7cb9_units_TryConnectAsSimple) new() any {
    return &p7e252d7cb.TryConnectAsSimple{}
}

func (inst* p7e252d7cb9_units_TryConnectAsSimple) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7e252d7cb.TryConnectAsSimple)
	nop(ie, com)

	
    com.EngineHost = inst.getEngineHost(ie)
    com.EnginePort = inst.getEnginePort(ie)
    com.RemoteService = inst.getRemoteService(ie)
    com.RemoteURL = inst.getRemoteURL(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p7e252d7cb9_units_TryConnectAsSimple) getEngineHost(ie application.InjectionExt)string{
    return ie.GetString("${debug.engine.host}")
}


func (inst*p7e252d7cb9_units_TryConnectAsSimple) getEnginePort(ie application.InjectionExt)int{
    return ie.GetInt("${debug.engine.port}")
}


func (inst*p7e252d7cb9_units_TryConnectAsSimple) getRemoteService(ie application.InjectionExt)string{
    return ie.GetString("${debug.case.connect-as-simple.service}")
}


func (inst*p7e252d7cb9_units_TryConnectAsSimple) getRemoteURL(ie application.InjectionExt)string{
    return ie.GetString("${debug.case.connect-as-simple.url}")
}


func (inst*p7e252d7cb9_units_TryConnectAsSimple) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.case.connect-as-simple.enabled}")
}



// type p7e252d7cb.Example in package:github.com/bitwormhole/what-sdk/src/debug/golang/units
//
// id:com-7e252d7cb97f6e6c-units-Example
// class:
// alias:
// scope:singleton
//
type p7e252d7cb9_units_Example struct {
}

func (inst* p7e252d7cb9_units_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7e252d7cb97f6e6c-units-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7e252d7cb9_units_Example) new() any {
    return &p7e252d7cb.Example{}
}

func (inst* p7e252d7cb9_units_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7e252d7cb.Example)
	nop(ie, com)

	


    return nil
}



// type p7e252d7cb.TryListenSimple in package:github.com/bitwormhole/what-sdk/src/debug/golang/units
//
// id:com-7e252d7cb97f6e6c-units-TryListenSimple
// class:
// alias:
// scope:singleton
//
type p7e252d7cb9_units_TryListenSimple struct {
}

func (inst* p7e252d7cb9_units_TryListenSimple) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7e252d7cb97f6e6c-units-TryListenSimple"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7e252d7cb9_units_TryListenSimple) new() any {
    return &p7e252d7cb.TryListenSimple{}
}

func (inst* p7e252d7cb9_units_TryListenSimple) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7e252d7cb.TryListenSimple)
	nop(ie, com)

	
    com.EngineHost = inst.getEngineHost(ie)
    com.EnginePort = inst.getEnginePort(ie)
    com.RemoteService = inst.getRemoteService(ie)
    com.RemoteURL = inst.getRemoteURL(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p7e252d7cb9_units_TryListenSimple) getEngineHost(ie application.InjectionExt)string{
    return ie.GetString("${debug.engine.host}")
}


func (inst*p7e252d7cb9_units_TryListenSimple) getEnginePort(ie application.InjectionExt)int{
    return ie.GetInt("${debug.engine.port}")
}


func (inst*p7e252d7cb9_units_TryListenSimple) getRemoteService(ie application.InjectionExt)string{
    return ie.GetString("${debug.case.listen-as-simple.service}")
}


func (inst*p7e252d7cb9_units_TryListenSimple) getRemoteURL(ie application.InjectionExt)string{
    return ie.GetString("${debug.case.listen-as-simple.url}")
}


func (inst*p7e252d7cb9_units_TryListenSimple) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.case.listen-as-simple.enabled}")
}


