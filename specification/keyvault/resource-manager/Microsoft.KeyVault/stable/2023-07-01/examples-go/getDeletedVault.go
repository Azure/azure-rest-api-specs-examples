package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/getDeletedVault.json
func ExampleVaultsClient_GetDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.DeletedVault = armkeyvault.DeletedVault{
	// 	Name: to.Ptr("sample-vault"),
	// 	Type: to.Ptr("Microsoft.KeyVault/deletedVaults"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedVaults/sample-vault"),
	// 	Properties: &armkeyvault.DeletedVaultProperties{
	// 		DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T00:00:59.000Z"); return t}()),
	// 		Location: to.Ptr("westus"),
	// 		PurgeProtectionEnabled: to.Ptr(true),
	// 		ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T00:00:59.000Z"); return t}()),
	// 		Tags: map[string]*string{
	// 		},
	// 		VaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
	// 	},
	// }
}
