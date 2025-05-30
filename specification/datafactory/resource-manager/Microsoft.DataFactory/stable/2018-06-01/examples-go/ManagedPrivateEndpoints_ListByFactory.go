package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_ListByFactory.json
func ExampleManagedPrivateEndpointsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedPrivateEndpointsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", nil)
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
		// page.ManagedPrivateEndpointListResponse = armdatafactory.ManagedPrivateEndpointListResponse{
		// 	Value: []*armdatafactory.ManagedPrivateEndpointResource{
		// 		{
		// 			Name: to.Ptr("exampleManagedPrivateEndpointName"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks/managedPrivateEndpoints"),
		// 			Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName/managedPrivateEndpoints/exampleManagedPrivateEndpointName"),
		// 			Properties: &armdatafactory.ManagedPrivateEndpoint{
		// 				ConnectionState: &armdatafactory.ConnectionStateProperties{
		// 					Description: to.Ptr(""),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Pending"),
		// 				},
		// 				Fqdns: []*string{
		// 				},
		// 				GroupID: to.Ptr("blob"),
		// 				PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
