package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBPrivateEndpointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rg1", "ddb1", "privateEndpointConnectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.PrivateEndpointConnectionsClientGetResponse{
	// 	PrivateEndpointConnection: &armcosmos.PrivateEndpointConnection{
	// 		Name: to.Ptr("privateEndpointConnectionName"),
	// 		Type: to.Ptr("Microsoft.DocumentDb/databaseAccounts/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDb/databaseAccounts/ddb1/privateEndpointConnections/privateEndpointConnectionName"),
	// 		Properties: &armcosmos.PrivateEndpointConnectionProperties{
	// 			GroupID: to.Ptr("Sql"),
	// 			PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
	// 				Description: to.Ptr("Auto-approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr("Approved"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
