package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBPrivateEndpointConnectionListGet.json
func ExamplePrivateEndpointConnectionsClient_NewListByDatabaseAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByDatabaseAccountPager("rg1", "ddb1", nil)
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
		// page = armcosmos.PrivateEndpointConnectionsClientListByDatabaseAccountResponse{
		// 	PrivateEndpointConnectionListResult: armcosmos.PrivateEndpointConnectionListResult{
		// 		Value: []*armcosmos.PrivateEndpointConnection{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDb/databaseAccounts/ddb1/privateEndpointConnections/privateEndpointConnectionName"),
		// 				Name: to.Ptr("privateEndpointConnectionName"),
		// 				Type: to.Ptr("Microsoft.DocumentDb/databaseAccounts/privateEndpointConnections"),
		// 				Properties: &armcosmos.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
		// 						Status: to.Ptr("Approved"),
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 					},
		// 					GroupID: to.Ptr("Sql"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDb/databaseAccounts/ddb1/privateEndpointConnections/privateEndpointConnectionName2"),
		// 				Name: to.Ptr("privateEndpointConnectionName"),
		// 				Type: to.Ptr("Microsoft.DocumentDb/databaseAccounts/privateEndpointConnections"),
		// 				Properties: &armcosmos.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName2"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
		// 						Status: to.Ptr("Approved"),
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 					},
		// 					GroupID: to.Ptr("Sql"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
