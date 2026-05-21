package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: 2023-03-02-preview/PrivateEndpointConnectionCreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "TestGroup", "testcluster", "private-endpoint-connection-name", armpostgresqlhsc.PrivateEndpointConnection{
		Properties: &armpostgresqlhsc.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armpostgresqlhsc.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr(armpostgresqlhsc.PrivateEndpointServiceConnectionStatusApproved),
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
	// res = armpostgresqlhsc.PrivateEndpointConnectionsClientCreateOrUpdateResponse{
	// 	PrivateEndpointConnection: &armpostgresqlhsc.PrivateEndpointConnection{
	// 		Name: to.Ptr("private-endpoint-connection-name"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateEndpointConnections/private-endpoint-connection-name"),
	// 		Properties: &armpostgresqlhsc.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armpostgresqlhsc.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armpostgresqlhsc.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armpostgresqlhsc.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armpostgresqlhsc.PrivateEndpointConnectionProvisioningState("Ready")),
	// 		},
	// 	},
	// }
}
