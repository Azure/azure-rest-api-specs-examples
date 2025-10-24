package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/listDeletedVaults.json
func ExampleVaultsClient_NewListDeletedPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultsClient().NewListDeletedPager(nil)
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
		// page = armkeyvault.VaultsClientListDeletedResponse{
		// 	DeletedVaultListResult: armkeyvault.DeletedVaultListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/deletedVaults?api-version=2025-05-01&%24skiptoken=HY3RaoMwAEX%2fRcbeYhJrnRXKYNWOuqpME0sfNcYui0Yxade19N8ne7hcDlzOvVuKX81eKKmt4G4dooLQwgqsL2NGHUDYV6o68Z4rY1e388RtNvRQn2vNJjEaMSgNvcbneMUcsKg8BFwft8DndQ0w9hu2QOiFLRs4TsNFNHzSMBFsGvTQGvuD%2f5bVuTOw4R03vPkH%2fVqNAlzm5SxfOwh7ACOA8POTlvPjILlaU1ke8jImOc23JCppQVfZnna0DXc4ISc3vSVuRo5zJE6%2bj25C3vwk2v2kEV2mMn7PyOc1DbtNGkonnzuLym1G400uI5QRZj0efw%3d%3d"),
		// 		Value: []*armkeyvault.DeletedVault{
		// 			{
		// 				Name: to.Ptr("vault-agile-drawer-6404"),
		// 				Type: to.Ptr("Microsoft.KeyVault/deletedVaults"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedVaults/sample-vault"),
		// 				Properties: &armkeyvault.DeletedVaultProperties{
		// 					DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T00:00:59Z"); return t}()),
		// 					Location: to.Ptr("westus"),
		// 					PurgeProtectionEnabled: to.Ptr(true),
		// 					ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T00:00:59Z"); return t}()),
		// 					Tags: map[string]*string{
		// 					},
		// 					VaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
