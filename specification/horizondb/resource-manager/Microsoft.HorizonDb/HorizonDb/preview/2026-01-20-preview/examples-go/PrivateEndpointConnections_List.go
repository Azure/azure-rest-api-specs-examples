package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/PrivateEndpointConnections_List.json
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
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("exampleresourcegroup", "examplecluster", nil)
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
		// 				Name: to.Ptr("exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/privateEndpointConnections/exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Properties: &armhorizondb.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("postgresqlServer"),
		// 					},
		// 					PrivateEndpoint: &armhorizondb.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningState("Ready")),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("exampleprivateendpointconnection2.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/privateEndpointConnections/exampleprivateendpointconnection2.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Properties: &armhorizondb.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("postgresqlServer"),
		// 					},
		// 					PrivateEndpoint: &armhorizondb.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningState("Ready")),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
