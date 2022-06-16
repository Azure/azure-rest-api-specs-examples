package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"ArunMonocle",
		"sdk-Namespace-2924",
		"privateEndpointConnectionName",
		armservicebus.PrivateEndpointConnection{
			Properties: &armservicebus.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armservicebus.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"),
				},
				PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
					Description: to.Ptr("testing"),
					Status:      to.Ptr(armservicebus.PrivateLinkConnectionStatusRejected),
				},
				ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
