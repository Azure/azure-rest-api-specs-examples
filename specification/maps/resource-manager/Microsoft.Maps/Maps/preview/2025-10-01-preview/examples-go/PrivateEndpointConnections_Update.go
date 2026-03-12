package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/PrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "myResourceGroup", "myMapsAccount", "privateEndpointConnectionName", armmaps.PrivateEndpointConnection{
		Properties: &armmaps.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armmaps.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armmaps.PrivateEndpointServiceConnectionStatusApproved),
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
	// res = armmaps.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnection: &armmaps.PrivateEndpointConnection{
	// 		Name: to.Ptr("privateEndpointConnectionName"),
	// 		Type: to.Ptr("Microsoft.Maps/accounts/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount/privateEndpointConnections/privateEndpointConnectionName"),
	// 		Properties: &armmaps.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armmaps.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armmaps.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Auto-Approved"),
	// 				Status: to.Ptr(armmaps.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armmaps.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
