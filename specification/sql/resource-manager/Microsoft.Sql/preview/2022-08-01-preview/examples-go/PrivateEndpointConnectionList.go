package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/PrivateEndpointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByServerPager("Default", "test-svr", nil)
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
		// page.PrivateEndpointConnectionListResult = armsql.PrivateEndpointConnectionListResult{
		// 	Value: []*armsql.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/test-svr/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 			Properties: &armsql.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armsql.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armsql.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr(armsql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 					Status: to.Ptr(armsql.PrivateLinkServiceConnectionStateStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armsql.PrivateEndpointProvisioningState("Succeeded")),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name-2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/test-svr/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 			Properties: &armsql.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armsql.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armsql.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr(armsql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 					Status: to.Ptr(armsql.PrivateLinkServiceConnectionStateStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armsql.PrivateEndpointProvisioningState("Succeeded")),
		// 			},
		// 	}},
		// }
	}
}
