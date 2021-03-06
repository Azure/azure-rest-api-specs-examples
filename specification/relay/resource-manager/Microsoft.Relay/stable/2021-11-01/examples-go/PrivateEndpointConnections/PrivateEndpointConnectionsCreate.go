package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/PrivateEndpointConnections/PrivateEndpointConnectionsCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrelay.NewPrivateEndpointConnectionsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourcegroup",
		"example-RelayNamespace-5849",
		"{privateEndpointConnection name}",
		armrelay.PrivateEndpointConnection{
			Properties: &armrelay.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armrelay.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"),
				},
				PrivateLinkServiceConnectionState: &armrelay.ConnectionState{
					Description: to.Ptr("You may pass"),
					Status:      to.Ptr(armrelay.PrivateLinkConnectionStatusApproved),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
