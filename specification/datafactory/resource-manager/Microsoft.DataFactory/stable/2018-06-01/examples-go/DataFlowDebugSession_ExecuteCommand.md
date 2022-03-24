Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_ExecuteCommand.json
func ExampleDataFlowDebugSessionClient_BeginExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewDataFlowDebugSessionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExecuteCommand(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.DataFlowDebugCommandRequest{
			Command: armdatafactory.DataFlowDebugCommandType("executePreviewQuery").ToPtr(),
			CommandPayload: &armdatafactory.DataFlowDebugCommandPayload{
				RowLimits:  to.Int32Ptr(100),
				StreamName: to.StringPtr("<stream-name>"),
			},
			SessionID: to.StringPtr("<session-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataFlowDebugSessionClientExecuteCommandResult)
}
```
