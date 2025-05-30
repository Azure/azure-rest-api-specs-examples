package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/PrivateEndpointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "Default", "test-svr", "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armpostgresqlflexibleservers.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/flexibleServers/test-svr/privateEndpointConnections/private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	Properties: &armpostgresqlflexibleservers.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("postgresqlServer")},
	// 			PrivateEndpoint: &armpostgresqlflexibleservers.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armpostgresqlflexibleservers.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Auto-approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armpostgresqlflexibleservers.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	}
}
