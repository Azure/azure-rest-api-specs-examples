package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.PrivateEndpointConnectionListResult = armpostgresql.PrivateEndpointConnectionListResult{
		// 	Value: []*armpostgresql.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/servers/test-svr/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 			Properties: &armpostgresql.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armpostgresql.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armpostgresql.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name-2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/servers/test-svr/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 			Properties: &armpostgresql.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armpostgresql.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armpostgresql.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
