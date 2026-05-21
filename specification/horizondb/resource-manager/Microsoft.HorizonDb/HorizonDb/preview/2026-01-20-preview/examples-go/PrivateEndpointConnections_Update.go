package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/PrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginUpdate(ctx, "exampleresourcegroup", "exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e", armhorizondb.PrivateEndpointConnectionUpdate{
		Properties: &armhorizondb.OptionalPropertiesUpdateableProperties{
			PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
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
	// res = armhorizondb.PrivateEndpointConnectionsClientUpdateResponse{
	// 	PrivateEndpointConnection: &armhorizondb.PrivateEndpointConnection{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/privateEndpointConnections/exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 		Name: to.Ptr("exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
	// 		Properties: &armhorizondb.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armhorizondb.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/examplePrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
	// 				Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
