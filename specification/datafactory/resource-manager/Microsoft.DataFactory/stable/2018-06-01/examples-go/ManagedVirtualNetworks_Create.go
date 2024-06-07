package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedVirtualNetworks_Create.json
func ExampleManagedVirtualNetworksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedVirtualNetworksClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", armdatafactory.ManagedVirtualNetworkResource{
		Properties: &armdatafactory.ManagedVirtualNetwork{},
	}, &armdatafactory.ManagedVirtualNetworksClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedVirtualNetworkResource = armdatafactory.ManagedVirtualNetworkResource{
	// 	Name: to.Ptr("exampleManagedVirtualNetworkName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName"),
	// 	Properties: &armdatafactory.ManagedVirtualNetwork{
	// 		Alias: to.Ptr("exampleFactoryName"),
	// 		VNetID: to.Ptr("12345678-1234-1234-1234-12345678123"),
	// 	},
	// }
}
