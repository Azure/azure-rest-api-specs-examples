package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
func ExampleDataFlowDebugSessionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowDebugSessionClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.CreateDataFlowDebugSessionRequest{
		IntegrationRuntime: &armdatafactory.IntegrationRuntimeDebugResource{
			Name: to.Ptr("ir1"),
			Properties: &armdatafactory.ManagedIntegrationRuntime{
				Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeManaged),
				TypeProperties: &armdatafactory.ManagedIntegrationRuntimeTypeProperties{
					ComputeProperties: &armdatafactory.IntegrationRuntimeComputeProperties{
						DataFlowProperties: &armdatafactory.IntegrationRuntimeDataFlowProperties{
							ComputeType: to.Ptr(armdatafactory.DataFlowComputeTypeGeneral),
							CoreCount:   to.Ptr[int32](48),
							TimeToLive:  to.Ptr[int32](10),
						},
						Location: to.Ptr("AutoResolve"),
					},
				},
			},
		},
		TimeToLive: to.Ptr[int32](60),
	}, nil)
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
