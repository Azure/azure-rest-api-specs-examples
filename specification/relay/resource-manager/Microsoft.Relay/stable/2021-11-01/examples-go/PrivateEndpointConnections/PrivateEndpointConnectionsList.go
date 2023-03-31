package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/PrivateEndpointConnections/PrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("myResourceGroup", "example-RelayNamespace-5849", nil)
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
		// page.PrivateEndpointConnectionListResult = armrelay.PrivateEndpointConnectionListResult{
		// 	Value: []*armrelay.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("{privateEndpointConnection name}"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces/PrivateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/alitest/providers/Microsoft.Relay/namespaces/relay-private-endpoint-test/privateEndpointConnections/{privateEndpointConnection name}"),
		// 			Location: to.Ptr("South Central US"),
		// 			Properties: &armrelay.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armrelay.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armrelay.ConnectionState{
		// 					Description: to.Ptr("approve please"),
		// 					Status: to.Ptr(armrelay.PrivateLinkConnectionStatusPending),
		// 				},
		// 				ProvisioningState: to.Ptr(armrelay.EndPointProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
