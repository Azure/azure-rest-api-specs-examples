package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/NameSpaces/PrivateEndPointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "SDK-ServiceBus-4794", "sdk-Namespace-5828", "privateEndpointConnectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.PrivateEndpointConnectionsClientGetResponse{
	// 	PrivateEndpointConnection: armservicebus.PrivateEndpointConnection{
	// 		Name: to.Ptr("privateEndpointConnectionName"),
	// 		Type: to.Ptr("Microsoft.ServiceBus/Namespaces/PrivateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5828/privateEndpointConnections/privateEndpointConnectionName"),
	// 		Properties: &armservicebus.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armservicebus.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5828"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
	// 				Description: to.Ptr("Auto-Approved"),
	// 				Status: to.Ptr(armservicebus.PrivateLinkConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
