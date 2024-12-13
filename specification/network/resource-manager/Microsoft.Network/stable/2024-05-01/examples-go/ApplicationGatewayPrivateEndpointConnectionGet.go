package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ApplicationGatewayPrivateEndpointConnectionGet.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().Get(ctx, "rg1", "appgw", "connection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGatewayPrivateEndpointConnection = armnetwork.ApplicationGatewayPrivateEndpointConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkResources/connection1"),
	// 	Name: to.Ptr("coonection1"),
	// 	Type: to.Ptr("Microsoft.Network/applicationGateways/privateEndpointConnections"),
	// 	Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
	// 		LinkIdentifier: to.Ptr("805319460"),
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approval Done"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
