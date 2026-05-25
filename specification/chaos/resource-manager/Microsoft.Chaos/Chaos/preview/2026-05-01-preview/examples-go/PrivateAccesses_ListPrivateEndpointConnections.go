package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/PrivateAccesses_ListPrivateEndpointConnections.json
func ExamplePrivateAccessesClient_NewListPrivateEndpointConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateAccessesClient().NewListPrivateEndpointConnectionsPager("myResourceGroup", "myPrivateAccess", nil)
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
		// page = armchaos.PrivateAccessesClientListPrivateEndpointConnectionsResponse{
		// 	PrivateEndpointConnectionListResult: armchaos.PrivateEndpointConnectionListResult{
		// 		Value: []*armchaos.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("myPrivateEndpointConnection"),
		// 				Type: to.Ptr("Microsoft.Chaos/privateAccesses/PrivateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess/privateEndpointConnections/myPrivateEndpointConnection"),
		// 				Properties: &armchaos.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armchaos.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armchaos.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armchaos.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
