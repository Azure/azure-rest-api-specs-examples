package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/privateEndpoint/PrivateEndpointConnection_List.json
func ExamplePrivateEndpointConnectionsClient_NewListByPrivateLinkScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByPrivateLinkScopePager("myResourceGroup", "myPrivateLinkScope", nil)
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
		// page.PrivateEndpointConnectionListResult = armhybridcompute.PrivateEndpointConnectionListResult{
		// 	Value: []*armhybridcompute.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 			Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
		// 				GroupIDs: []*string{
		// 					to.Ptr("hybridcompute")},
		// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
		// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr("Approved"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("private-endpoint-connection-name-2"),
		// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("hybridcompute")},
		// 						PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
		// 							ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 						},
		// 						PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
		// 							Description: to.Ptr("Please approve my connection."),
		// 							ActionsRequired: to.Ptr("None"),
		// 							Status: to.Ptr("Pending"),
		// 						},
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 					},
		// 			}},
		// 		}
	}
}
