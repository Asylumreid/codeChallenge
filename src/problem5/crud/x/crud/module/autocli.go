package crud

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "crud/api/crud/crud"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowPost",
					Use:            "show-post [id]",
					Short:          "Query show-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListPost",
					Use:            "list-post",
					Short:          "Query list-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListSpecificPosts",
					Use:            "list-specific-posts [event-type]",
					Short:          "Query list-specific-posts",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "eventType"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateEventPost",
					Use:            "create-event-post [title] [event-type] [datetime] [venue] [description] [interest]",
					Short:          "Send a create-event-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "eventType"}, {ProtoField: "datetime"}, {ProtoField: "venue"}, {ProtoField: "description"}, {ProtoField: "interest"}},
				},
				{
					RpcMethod:      "UpdateEventPost",
					Use:            "update-event-post [title] [event-type] [datetime] [venue] [description] [interest] [id]",
					Short:          "Send a update-event-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "eventType"}, {ProtoField: "datetime"}, {ProtoField: "venue"}, {ProtoField: "description"}, {ProtoField: "interest"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteEventPost",
					Use:            "delete-event-post [id]",
					Short:          "Send a delete-event-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ExpressInterest",
					Use:            "express-interest [id]",
					Short:          "Send a express-interest tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
