package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyPrivateLinkScope", "private-endpoint-connection-name", armmonitor.PrivateEndpointConnection{
		Properties: &armmonitor.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armmonitor.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr(armmonitor.PrivateEndpointServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnection = armmonitor.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.Insights/privateLinkScopes/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armmonitor.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armmonitor.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armmonitor.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armmonitor.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armmonitor.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
