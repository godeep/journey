package plugins

import (
	"github.com/kabukky/journey/structure"
	"github.com/yuin/gopher-lua"
	"log"
)

func Execute(name string, values *structure.RequestData) ([]byte, error) {
	// Retrieve a lua state
	vm := values.PluginVMs[name]
	// Execute plugin
	err := vm.CallByParam(lua.P{Fn: vm.GetGlobal(name), NRet: 1, Protect: true})
	if err != nil {
		log.Println("Error while exec:", err)
		// Since the vm threw an error, close it and don't put the map back into the pool
		for _, luavm := range values.PluginVMs {
			luavm.Close()
		}
		values.PluginVMs = nil
		return []byte{}, err
	}
	ret := vm.ToString(-1)
	return []byte(ret), nil
}