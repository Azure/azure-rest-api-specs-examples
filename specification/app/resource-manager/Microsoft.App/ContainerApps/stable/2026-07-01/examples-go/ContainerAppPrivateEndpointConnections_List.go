package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/ContainerAppPrivateEndpointConnections_List.json
func ExampleContainerAppPrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("8efdecc5-919e-44eb-b179-915dca89ebf9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppPrivateEndpointConnectionsClient().NewListPager("examplerg", "testcontainerapp0", nil)
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
		// page = armappcontainers.ContainerAppPrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionListResult: armappcontainers.PrivateEndpointConnectionListResult{
		// 		Value: []*armappcontainers.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("test-private-endpoint-connection"),
		// 				Type: to.Ptr("Microsoft.App/containerApps/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/containerApps/testcontainerapp0/privateEndpointConnections/test-private-endpoint-connection"),
		// 				Properties: &armappcontainers.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("containerApps"),
		// 					},
		// 					PrivateEndpoint: &armappcontainers.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Network/privateEndpoints/test-private-endpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armappcontainers.PrivateLinkServiceConnectionState{
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armappcontainers.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armappcontainers.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
