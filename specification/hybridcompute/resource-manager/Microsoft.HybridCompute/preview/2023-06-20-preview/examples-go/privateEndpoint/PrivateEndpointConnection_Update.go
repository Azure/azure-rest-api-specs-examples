package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateEndpoint/PrivateEndpointConnection_Update.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myPrivateLinkScope", "private-endpoint-connection-name", armhybridcompute.PrivateEndpointConnection{
		Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr("Approved"),
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
	// res.PrivateEndpointConnection = armhybridcompute.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("hybridcompute")},
	// 			PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 				ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 				Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr("Approved"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}
