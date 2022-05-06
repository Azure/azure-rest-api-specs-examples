Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.5.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
func ExampleDataFlowDebugSessionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowDebugSessionClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.CreateDataFlowDebugSessionRequest{
			IntegrationRuntime: &armdatafactory.IntegrationRuntimeDebugResource{
				Name: to.Ptr("<name>"),
				Properties: &armdatafactory.ManagedIntegrationRuntime{
					Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeManaged),
					TypeProperties: &armdatafactory.ManagedIntegrationRuntimeTypeProperties{
						ComputeProperties: &armdatafactory.IntegrationRuntimeComputeProperties{
							DataFlowProperties: &armdatafactory.IntegrationRuntimeDataFlowProperties{
								ComputeType: to.Ptr(armdatafactory.DataFlowComputeTypeGeneral),
								CoreCount:   to.Ptr[int32](48),
								TimeToLive:  to.Ptr[int32](10),
							},
							Location: to.Ptr("<location>"),
						},
					},
				},
			},
			TimeToLive: to.Ptr[int32](60),
		},
		&armdatafactory.DataFlowDebugSessionClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
