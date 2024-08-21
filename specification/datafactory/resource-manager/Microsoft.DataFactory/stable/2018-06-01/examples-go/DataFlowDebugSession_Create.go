package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
func ExampleDataFlowDebugSessionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataFlowDebugSessionClient().BeginCreate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.CreateDataFlowDebugSessionRequest{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CreateDataFlowDebugSessionResponse = armdatafactory.CreateDataFlowDebugSessionResponse{
	// 	SessionID: to.Ptr("229c688c-944c-44ac-b31a-82d50f347154"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
