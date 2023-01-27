package app_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	cosmossecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	quicksilverapp "github.com/ingenuity-build/quicksilver/app"
	"github.com/ingenuity-build/quicksilver/app/testutil"
)

func init() {
	simapp.GetSimulatorFlags()
}

func TestQuicksilverExport(t *testing.T) {
	privVal := mock.NewPV()
	pubKey, _ := privVal.GetPubKey()

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	senderPubKey := cosmossecp256k1.PubKey{
		Key: senderPrivKey.PubKey().Bytes(),
	}

	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), &senderPubKey, 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}
	db := dbm.NewMemDB()
	app := quicksilverapp.NewQuicksilver(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
		db,
		nil,
		true,
		map[int64]bool{},
		quicksilverapp.DefaultNodeHome,
		0,
		quicksilverapp.MakeEncodingConfig(),
		wasm.EnableAllProposals,
		simapp.EmptyAppOptions{},
		quicksilverapp.GetWasmOpts(simapp.EmptyAppOptions{}),
		false,
	)

	genesisState := quicksilverapp.NewDefaultGenesisState()
	genesisState = testutil.GenesisStateWithValSet(t, app, genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)

	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			ChainId:       "quicksilver-1",
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := quicksilverapp.NewQuicksilver(log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
		db,
		nil,
		true,
		map[int64]bool{},
		quicksilverapp.DefaultNodeHome,
		0,
		quicksilverapp.MakeEncodingConfig(),
		wasm.EnableAllProposals,
		simapp.EmptyAppOptions{},
		quicksilverapp.GetWasmOpts(simapp.EmptyAppOptions{}),
		false,
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestAppStateDeterminism(t *testing.T) {
	if !simapp.FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := simapp.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = true
	config.AllInvariants = true
	config.NumBlocks = 100
	config.BlockSize = 200
	config.Commit = true

	numSeeds := 3
	numTimesToRunPerSeed := 5
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			db := dbm.NewMemDB()
			app := quicksilverapp.NewQuicksilver(
				log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
				db,
				nil,
				true,
				map[int64]bool{},
				quicksilverapp.DefaultNodeHome,
				0,
				quicksilverapp.MakeEncodingConfig(),
				wasm.EnableAllProposals,
				simapp.EmptyAppOptions{},
				quicksilverapp.GetWasmOpts(simapp.EmptyAppOptions{}),
				false,
			)

			fmt.Printf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				app.GetBaseApp(),
				testutil.AppStateFn(app.AppCodec(), app.SimulationManager()),
				simulationtypes.RandomAccounts,
				simapp.SimulationOperations(app, app.AppCodec(), config),
				app.ModuleAccountAddrs(),
				config,
				app.AppCodec(),
			)
			require.NoError(t, err)

			if config.Commit {
				simapp.PrintStats(db)
			}

			appHash := app.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
