package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/PrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "exampleresourcegroup", "exampleserver", "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.PrivateEndpointConnectionsClientGetResponse{
	// 	PrivateEndpointConnection: &armpostgresqlflexibleservers.PrivateEndpointConnection{
	// 		Name: to.Ptr("private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/privateEndpointConnections/private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 		Properties: &armpostgresqlflexibleservers.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("postgresqlServer"),
	// 			},
	// 			PrivateEndpoint: &armpostgresqlflexibleservers.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armpostgresqlflexibleservers.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Auto-approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointConnectionProvisioningState("Ready")),
	// 		},
	// 	},
	// }
}
