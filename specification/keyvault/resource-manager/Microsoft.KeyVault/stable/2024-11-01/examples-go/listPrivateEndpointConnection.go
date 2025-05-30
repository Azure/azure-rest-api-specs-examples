package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/listPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByResourcePager("sample-group", "sample-vault", nil)
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
		// page.PrivateEndpointConnectionListResult = armkeyvault.PrivateEndpointConnectionListResult{
		// 	Value: []*armkeyvault.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("sample-pec"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/privateEndpointConnections/sample-pec"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-pec"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/privateEndpointConnections/sample-pec"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
