package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2026-05-15/getDeletedVault.json
func ExampleVaultsClient_GetDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVaultsClient().GetDeleted(ctx, "sample-vault", "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.VaultsClientGetDeletedResponse{
	// 	DeletedVault: armkeyvault.DeletedVault{
	// 		Name: to.Ptr("sample-vault"),
	// 		Type: to.Ptr("Microsoft.KeyVault/deletedVaults"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedVaults/sample-vault"),
	// 		Properties: &armkeyvault.DeletedVaultProperties{
	// 			DeletionDate: to.Ptr(time.Date(2017, time.January, 1, 0, 0, 59, 0, time.UTC)),
	// 			Location: to.Ptr("westus"),
	// 			PurgeProtectionEnabled: to.Ptr(true),
	// 			ScheduledPurgeDate: to.Ptr(time.Date(2017, time.April, 1, 0, 0, 59, 0, time.UTC)),
	// 			Tags: map[string]*string{
	// 			},
	// 			VaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
	// 		},
	// 	},
	// }
}
