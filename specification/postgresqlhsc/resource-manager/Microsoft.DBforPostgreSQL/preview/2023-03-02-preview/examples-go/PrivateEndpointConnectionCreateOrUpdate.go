package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/PrivateEndpointConnectionCreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "TestGroup", "testcluster", "private-endpoint-connection-name", armcosmosforpostgresql.PrivateEndpointConnection{
		Properties: &armcosmosforpostgresql.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcosmosforpostgresql.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr(armcosmosforpostgresql.PrivateEndpointServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnection = armcosmosforpostgresql.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armcosmosforpostgresql.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcosmosforpostgresql.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcosmosforpostgresql.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armcosmosforpostgresql.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armcosmosforpostgresql.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
