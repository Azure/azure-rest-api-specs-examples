package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/ManagedVirtualNetworks_ListByFactory.json
func ExampleManagedVirtualNetworksClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedVirtualNetworksClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdatafactory.ManagedVirtualNetworksClientListByFactoryResponse{
		// 	ManagedVirtualNetworkListResponse: armdatafactory.ManagedVirtualNetworkListResponse{
		// 		Value: []*armdatafactory.ManagedVirtualNetworkResource{
		// 			{
		// 				Name: to.Ptr("exampleManagedVirtualNetworkName"),
		// 				Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks"),
		// 				Etag: to.Ptr("0400f1a1-0000-0000-0000-5b2188640000"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName"),
		// 				Properties: &armdatafactory.ManagedVirtualNetwork{
		// 					Alias: to.Ptr("exampleFactoryName"),
		// 					VNetID: to.Ptr("5a7bd944-87e6-454a-8d4d-9fba446514fd"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
