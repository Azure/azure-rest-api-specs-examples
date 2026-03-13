package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/PrivateEndpointConnections_ListByAccount.json
func ExamplePrivateEndpointConnectionsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByAccountPager("myResourceGroup", "myMapsAccount", nil)
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
		// page = armmaps.PrivateEndpointConnectionsClientListByAccountResponse{
		// 	PrivateEndpointConnectionList: armmaps.PrivateEndpointConnectionList{
		// 		Value: []*armmaps.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("privateEndpointConnectionName"),
		// 				Type: to.Ptr("Microsoft.Maps/accounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount/privateEndpointConnections/privateEndpointConnectionName"),
		// 				Properties: &armmaps.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armmaps.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armmaps.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						Status: to.Ptr(armmaps.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armmaps.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("privateEndpointConnectionName2"),
		// 				Type: to.Ptr("Microsoft.Maps/accounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount/privateEndpointConnections/privateEndpointConnectionName2"),
		// 				Properties: &armmaps.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armmaps.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/petest02"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armmaps.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						Status: to.Ptr(armmaps.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armmaps.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
