package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/PrivateEndpointConnections_ApproveOrReject.json
func ExamplePrivateEndpointConnectionsClient_BeginApproveOrReject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginApproveOrReject(ctx, "example-resource-group", "example-cluster", "example-private-endpoint-connection", armhorizondb.PrivateEndpointConnectionResource{
		Properties: &armhorizondb.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
				Status:      to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
				Description: to.Ptr("Approved for the production database network"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.PrivateEndpointConnectionsClientApproveOrRejectResponse{
	// 	PrivateEndpointConnectionResource: armhorizondb.PrivateEndpointConnectionResource{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/privateEndpointConnections/example-private-endpoint-connection"),
	// 		Name: to.Ptr("example-private-endpoint-connection"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/privateEndpointConnections"),
	// 		Properties: &armhorizondb.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armhorizondb.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/network-resource-group/providers/Microsoft.Network/privateEndpoints/example-private-endpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhorizondb.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armhorizondb.PrivateEndpointServiceConnectionStatusApproved),
	// 				Description: to.Ptr("Approved for the production database network"),
	// 				ActionsRequired: to.Ptr("None"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhorizondb.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
