package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().CreateOrUpdate(ctx, "ArunMonocle", "sdk-Namespace-2924", "privateEndpointConnectionName", armservicebus.PrivateEndpointConnection{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armservicebus.PrivateEndpointConnection{
	// 	Name: to.Ptr("928c44d5-b7c6-423b-b6fa-811e0c27b3e0"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5828/privateEndpointConnections/928c44d5-b7c6-423b-b6fa-811e0c27b3e0"),
	// 	Properties: &armservicebus.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armservicebus.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5828"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			Status: to.Ptr(armservicebus.PrivateLinkConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
	// 	},
	// }
}
