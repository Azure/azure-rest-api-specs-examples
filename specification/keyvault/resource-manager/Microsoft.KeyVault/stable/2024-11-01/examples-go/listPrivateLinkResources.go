package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/listPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_ListByVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByVault(ctx, "sample-group", "sample-vault", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armkeyvault.PrivateLinkResourceListResult{
	// 	Value: []*armkeyvault.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("vault"),
	// 			Type: to.Ptr("Microsoft.KeyVault/vaults/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault/privateLinkResources/vault"),
	// 			Properties: &armkeyvault.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("vault"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("default")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.vaultcore.azure.net")},
	// 					},
	// 			}},
	// 		}
}
