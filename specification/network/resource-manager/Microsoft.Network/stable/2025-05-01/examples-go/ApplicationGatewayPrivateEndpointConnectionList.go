package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ApplicationGatewayPrivateEndpointConnectionList.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().NewListPager("rg1", "appgw", nil)
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
		// page.ApplicationGatewayPrivateEndpointConnectionListResult = armnetwork.ApplicationGatewayPrivateEndpointConnectionListResult{
		// 	Value: []*armnetwork.ApplicationGatewayPrivateEndpointConnection{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkResources/connection1"),
		// 			Name: to.Ptr("coonection1"),
		// 			Type: to.Ptr("Microsoft.Network/applicationGateways/privateEndpointConnections"),
		// 			Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
		// 				LinkIdentifier: to.Ptr("805319460"),
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Approval Done"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
