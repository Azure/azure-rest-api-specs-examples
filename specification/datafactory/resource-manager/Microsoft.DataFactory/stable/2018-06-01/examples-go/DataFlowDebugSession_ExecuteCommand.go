package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/69ece3818b8b0929b43a07c3fe25716427734882/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_ExecuteCommand.json
func ExampleDataFlowDebugSessionClient_BeginExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataFlowDebugSessionClient().BeginExecuteCommand(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DataFlowDebugCommandRequest{
		Command: to.Ptr(armdatafactory.DataFlowDebugCommandTypeExecutePreviewQuery),
		CommandPayload: &armdatafactory.DataFlowDebugCommandPayload{
			RowLimits:  to.Ptr[int32](100),
			StreamName: to.Ptr("source1"),
		},
		SessionID: to.Ptr("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowDebugCommandResponse = armdatafactory.DataFlowDebugCommandResponse{
	// 	Data: to.Ptr("some output"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
