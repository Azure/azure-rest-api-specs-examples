package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/69ece3818b8b0929b43a07c3fe25716427734882/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Start.json
func ExampleIntegrationRuntimesClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginStart(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedIntegrationRuntime", nil)
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
	// res.IntegrationRuntimeStatusResponse = armdatafactory.IntegrationRuntimeStatusResponse{
	// 	Name: to.Ptr("exampleManagedIntegrationRuntime"),
	// 	Properties: &armdatafactory.ManagedIntegrationRuntimeStatus{
	// 		Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeManaged),
	// 		DataFactoryName: to.Ptr("exampleFactoryName"),
	// 		State: to.Ptr(armdatafactory.IntegrationRuntimeStateStarted),
	// 		TypeProperties: &armdatafactory.ManagedIntegrationRuntimeStatusTypeProperties{
	// 			CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T21:11:01.869Z"); return t}()),
	// 			Nodes: []*armdatafactory.ManagedIntegrationRuntimeNode{
	// 			},
	// 			OtherErrors: []*armdatafactory.ManagedIntegrationRuntimeError{
	// 			},
	// 		},
	// 	},
	// }
}
