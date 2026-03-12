package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview/v2"
)

// Generated from example definition: 2024-04-01-preview/IngestionPrivateEndpointConnections_List.json
func ExampleIngestionPrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIngestionPrivateEndpointConnectionsClient().NewListPager("SampleResourceGroup", "account1", nil)
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
		// page = armpurview.IngestionPrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionList: armpurview.PrivateEndpointConnectionList{
		// 		Value: []*armpurview.PrivateEndpointConnection{
		// 			{
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/ingestionPrivateEndpointConnections/privateEndpointConnection1"),
		// 				Properties: &armpurview.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armpurview.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint1"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Approved by johndoe@company.com"),
		// 						Status: to.Ptr(armpurview.PrivateEndpointConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 			{
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/ingestionPrivateEndpointConnections/privateEndpointConnection2"),
		// 				Properties: &armpurview.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armpurview.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Rejected by johndoe@company.com"),
		// 						Status: to.Ptr(armpurview.PrivateEndpointConnectionStatusRejected),
		// 					},
		// 					ProvisioningState: to.Ptr("Deleting"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
