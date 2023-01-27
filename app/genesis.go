package app

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	encCfg := MakeEncodingConfig()
	gen := ModuleBasics.DefaultGenesis(encCfg.Marshaler)

	// here we override wasm config to make it permissioned by default
	wasmGen := wasm.GenesisState{
		Params: wasmtypes.Params{
			CodeUploadAccess:             wasmtypes.AllowNobody,
			InstantiateDefaultPermission: wasmtypes.AccessTypeEverybody,
		},
	}
	gen[wasm.ModuleName] = encCfg.Marshaler.MustMarshalJSON(&wasmGen)
	return gen
}
