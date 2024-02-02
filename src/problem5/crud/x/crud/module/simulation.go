package crud

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"crud/testutil/sample"
	crudsimulation "crud/x/crud/simulation"
	"crud/x/crud/types"
)

// avoid unused import issue
var (
	_ = crudsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateEventPost = "op_weight_msg_create_event_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEventPost int = 100

	opWeightMsgUpdateEventPost = "op_weight_msg_update_event_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEventPost int = 100

	opWeightMsgDeleteEventPost = "op_weight_msg_delete_event_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEventPost int = 100

	opWeightMsgExpressInterest = "op_weight_msg_express_interest"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExpressInterest int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	crudGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&crudGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEventPost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateEventPost, &weightMsgCreateEventPost, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEventPost = defaultWeightMsgCreateEventPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEventPost,
		crudsimulation.SimulateMsgCreateEventPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEventPost int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateEventPost, &weightMsgUpdateEventPost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEventPost = defaultWeightMsgUpdateEventPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEventPost,
		crudsimulation.SimulateMsgUpdateEventPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEventPost int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteEventPost, &weightMsgDeleteEventPost, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEventPost = defaultWeightMsgDeleteEventPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEventPost,
		crudsimulation.SimulateMsgDeleteEventPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExpressInterest int
	simState.AppParams.GetOrGenerate(opWeightMsgExpressInterest, &weightMsgExpressInterest, nil,
		func(_ *rand.Rand) {
			weightMsgExpressInterest = defaultWeightMsgExpressInterest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExpressInterest,
		crudsimulation.SimulateMsgExpressInterest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEventPost,
			defaultWeightMsgCreateEventPost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudsimulation.SimulateMsgCreateEventPost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEventPost,
			defaultWeightMsgUpdateEventPost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudsimulation.SimulateMsgUpdateEventPost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEventPost,
			defaultWeightMsgDeleteEventPost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudsimulation.SimulateMsgDeleteEventPost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgExpressInterest,
			defaultWeightMsgExpressInterest,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudsimulation.SimulateMsgExpressInterest(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
