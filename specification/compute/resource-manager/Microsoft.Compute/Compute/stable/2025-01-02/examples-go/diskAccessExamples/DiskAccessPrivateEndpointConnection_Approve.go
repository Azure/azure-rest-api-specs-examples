package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/diskAccessExamples/DiskAccessPrivateEndpointConnection_Approve.json
func ExampleDiskAccessesClient_BeginUpdateAPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskAccessesClient().BeginUpdateAPrivateEndpointConnection(ctx, "myResourceGroup", "myDiskAccess", "myPrivateEndpointConnection", armcompute.PrivateEndpointConnection{
		Properties: &armcompute.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcompute.PrivateLinkServiceConnectionState{
				Status:      to.Ptr(armcompute.PrivateEndpointServiceConnectionStatusApproved),
				Description: to.Ptr("Approving myPrivateEndpointConnection"),
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
	// res = armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse{
	// 	PrivateEndpointConnection: armcompute.PrivateEndpointConnection{
	// 		Name: to.Ptr("myPrivateEndpointConnectionName"),
	// 		Type: to.Ptr("Microsoft.Compute/diskAccesses/PrivateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess/privateEndpoinConnections/myPrivateEndpointConnectionName"),
	// 		Properties: &armcompute.PrivateEndpointConnectionProperties{
	// 			ProvisioningState: to.Ptr(armcompute.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			PrivateEndpoint: &armcompute.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armcompute.PrivateLinkServiceConnectionState{
	// 				ActionsRequired: to.Ptr("None"),
	// 				Description: to.Ptr("Approving myPrivateEndpointConnection"),
	// 				Status: to.Ptr(armcompute.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 		},
	// 	},
	// }
}
