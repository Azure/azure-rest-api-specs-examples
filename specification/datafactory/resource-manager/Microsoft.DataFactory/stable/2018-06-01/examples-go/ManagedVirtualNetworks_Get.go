package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedVirtualNetworks_Get.json
func ExampleManagedVirtualNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedVirtualNetworksClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", &armdatafactory.ManagedVirtualNetworksClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedVirtualNetworkResource = armdatafactory.ManagedVirtualNetworkResource{
	// 	Name: to.Ptr("exampleManagedVirtualNetworkName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks"),
	// 	Etag: to.Ptr("15003c4f-0000-0200-0000-5cbe090b0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName"),
	// 	Properties: &armdatafactory.ManagedVirtualNetwork{
	// 		Alias: to.Ptr("exampleFactoryName"),
	// 		VNetID: to.Ptr("5a7bd944-87e6-454a-8d4d-9fba446514fd"),
	// 	},
	// }
}
