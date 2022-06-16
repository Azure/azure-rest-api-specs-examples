package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_ExecuteCommand.json
func ExampleDataFlowDebugSessionClient_BeginExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowDebugSessionClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExecuteCommand(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		armdatafactory.DataFlowDebugCommandRequest{
			Command: to.Ptr(armdatafactory.DataFlowDebugCommandTypeExecutePreviewQuery),
			CommandPayload: &armdatafactory.DataFlowDebugCommandPayload{
				RowLimits:  to.Ptr[int32](100),
				StreamName: to.Ptr("source1"),
			},
			SessionID: to.Ptr("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
