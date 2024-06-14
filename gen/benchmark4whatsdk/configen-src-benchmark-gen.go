package benchmark4whatsdk
import (
    pb6853deb0 "github.com/bitwormhole/what-sdk/src/benchmark/golang/bench"
     "github.com/starter-go/application"
)

// type pb6853deb0.ProxyAgent in package:github.com/bitwormhole/what-sdk/src/benchmark/golang/bench
//
// id:com-b6853deb00f44e8c-bench-ProxyAgent
// class:
// alias:alias-b6853deb00f44e8c336028d971848c80-Agent
// scope:singleton
//
type pb6853deb00_bench_ProxyAgent struct {
}

func (inst* pb6853deb00_bench_ProxyAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b6853deb00f44e8c-bench-ProxyAgent"
	r.Classes = ""
	r.Aliases = "alias-b6853deb00f44e8c336028d971848c80-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb6853deb00_bench_ProxyAgent) new() any {
    return &pb6853deb0.ProxyAgent{}
}

func (inst* pb6853deb00_bench_ProxyAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb6853deb0.ProxyAgent)
	nop(ie, com)

	
    com.Host1 = inst.getHost1(ie)
    com.Port1 = inst.getPort1(ie)
    com.Host2 = inst.getHost2(ie)
    com.Port2 = inst.getPort2(ie)


    return nil
}


func (inst*pb6853deb00_bench_ProxyAgent) getHost1(ie application.InjectionExt)string{
    return ie.GetString("${benchmark.core1.host}")
}


func (inst*pb6853deb00_bench_ProxyAgent) getPort1(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.core1.port}")
}


func (inst*pb6853deb00_bench_ProxyAgent) getHost2(ie application.InjectionExt)string{
    return ie.GetString("${benchmark.core2.host}")
}


func (inst*pb6853deb00_bench_ProxyAgent) getPort2(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.core2.port}")
}



// type pb6853deb0.Client in package:github.com/bitwormhole/what-sdk/src/benchmark/golang/bench
//
// id:com-b6853deb00f44e8c-bench-Client
// class:
// alias:
// scope:singleton
//
type pb6853deb00_bench_Client struct {
}

func (inst* pb6853deb00_bench_Client) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b6853deb00f44e8c-bench-Client"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb6853deb00_bench_Client) new() any {
    return &pb6853deb0.Client{}
}

func (inst* pb6853deb00_bench_Client) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb6853deb0.Client)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ServiceName = inst.getServiceName(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Interval = inst.getInterval(ie)
    com.RequestX = inst.getRequestX(ie)
    com.ResponseX = inst.getResponseX(ie)
    com.PayloadSize = inst.getPayloadSize(ie)
    com.RemoteURL = inst.getRemoteURL(ie)


    return nil
}


func (inst*pb6853deb00_bench_Client) getAgent(ie application.InjectionExt)pb6853deb0.Agent{
    return ie.GetComponent("#alias-b6853deb00f44e8c336028d971848c80-Agent").(pb6853deb0.Agent)
}


func (inst*pb6853deb00_bench_Client) getServiceName(ie application.InjectionExt)string{
    return ie.GetString("${benchmark.client.service}")
}


func (inst*pb6853deb00_bench_Client) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${benchmark.client.enabled}")
}


func (inst*pb6853deb00_bench_Client) getInterval(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.client.interval}")
}


func (inst*pb6853deb00_bench_Client) getRequestX(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.client.request-x}")
}


func (inst*pb6853deb00_bench_Client) getResponseX(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.client.response-x}")
}


func (inst*pb6853deb00_bench_Client) getPayloadSize(ie application.InjectionExt)int{
    return ie.GetInt("${benchmark.client.payload-size}")
}


func (inst*pb6853deb00_bench_Client) getRemoteURL(ie application.InjectionExt)string{
    return ie.GetString("${benchmark.client.remote-url}")
}



// type pb6853deb0.Server in package:github.com/bitwormhole/what-sdk/src/benchmark/golang/bench
//
// id:com-b6853deb00f44e8c-bench-Server
// class:
// alias:
// scope:singleton
//
type pb6853deb00_bench_Server struct {
}

func (inst* pb6853deb00_bench_Server) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b6853deb00f44e8c-bench-Server"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb6853deb00_bench_Server) new() any {
    return &pb6853deb0.Server{}
}

func (inst* pb6853deb00_bench_Server) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb6853deb0.Server)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ServiceName = inst.getServiceName(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*pb6853deb00_bench_Server) getAgent(ie application.InjectionExt)pb6853deb0.Agent{
    return ie.GetComponent("#alias-b6853deb00f44e8c336028d971848c80-Agent").(pb6853deb0.Agent)
}


func (inst*pb6853deb00_bench_Server) getServiceName(ie application.InjectionExt)string{
    return ie.GetString("${benchmark.server.service}")
}


func (inst*pb6853deb00_bench_Server) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${benchmark.server.enabled}")
}


