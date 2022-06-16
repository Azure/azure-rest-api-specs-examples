package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Create.json
func ExampleManagedPrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewManagedPrivateEndpointsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"exampleManagedVirtualNetworkName",
		"exampleManagedPrivateEndpointName",
		armdatafactory.ManagedPrivateEndpointResource{
			Properties: &armdatafactory.ManagedPrivateEndpoint{
				Fqdns:                 []*string{},
				GroupID:               to.Ptr("blob"),
				PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
			},
		},
		&armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
