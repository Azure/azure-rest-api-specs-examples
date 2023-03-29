package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstancePrivateEndpointConnectionUpdate.json
func ExampleManagedInstancePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstancePrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "Default", "test-cl", "private-endpoint-connection-name", armsql.ManagedInstancePrivateEndpointConnection{
		Properties: &armsql.ManagedInstancePrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armsql.ManagedInstancePrivateLinkServiceConnectionStateProperty{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr("Approved"),
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
	// res.ManagedInstancePrivateEndpointConnection = armsql.ManagedInstancePrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/test-cl/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armsql.ManagedInstancePrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armsql.ManagedInstancePrivateEndpointProperty{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armsql.ManagedInstancePrivateLinkServiceConnectionStateProperty{
	// 			Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
