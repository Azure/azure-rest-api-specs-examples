package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/PrivateEndpointConnections_List.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("example-resource-group", "example-cluster", nil)
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
		// page = armhorizondb.PrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionResourceListResult: armhorizondb.PrivateEndpointConnectionResourceListResult{
		// 		Value: []*armhorizondb.PrivateEndpointConnectionResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/privateEndpointConnections/example-private-endpoint-connection"),
		// 				Name: to.Ptr("example-private-endpoint-connection"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
		// 				Properties: &armhorizondb.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("example-cluster"),
		// 					},
		// 					PrivateEndpoint: &armhorizondb.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/network-resource-group/providers/Microsoft.Network/privateEndpoints/example-private-endpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
		// 						Description: to.Ptr("Approved for the production database network"),
		// 						ActionsRequired: to.Ptr("None"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/privateEndpointConnections/example-pool-private-endpoint-connection"),
		// 				Name: to.Ptr("example-pool-private-endpoint-connection"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
		// 				Properties: &armhorizondb.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("example-pool"),
		// 					},
		// 					PrivateEndpoint: &armhorizondb.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/network-resource-group/providers/Microsoft.Network/privateEndpoints/example-pool-private-endpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
		// 						Description: to.Ptr("Approved for read-only Pool access"),
		// 						ActionsRequired: to.Ptr("None"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
