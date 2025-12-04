package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/PrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByServerPager("exampleresourcegroup", "exampleserver", nil)
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
		// page.PrivateEndpointConnectionList = armpostgresqlflexibleservers.PrivateEndpointConnectionList{
		// 	Value: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/privateEndpointConnections/private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 			Properties: &armpostgresqlflexibleservers.PrivateEndpointConnectionProperties{
		// 				GroupIDs: []*string{
		// 					to.Ptr("postgresqlServer")},
		// 					PrivateEndpoint: &armpostgresqlflexibleservers.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armpostgresqlflexibleservers.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("private-endpoint-connection-name-2.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/privateEndpointConnections/private-endpoint-connection-name-2.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
		// 				Properties: &armpostgresqlflexibleservers.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("postgresqlServer")},
		// 						PrivateEndpoint: &armpostgresqlflexibleservers.PrivateEndpoint{
		// 							ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 						},
		// 						PrivateLinkServiceConnectionState: &armpostgresqlflexibleservers.PrivateLinkServiceConnectionState{
		// 							Description: to.Ptr("Auto-approved"),
		// 							ActionsRequired: to.Ptr("None"),
		// 							Status: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointServiceConnectionStatusApproved),
		// 						},
		// 						ProvisioningState: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}
