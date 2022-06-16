package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/VaultCRUD/PutBackupVault.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupVaultsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armdataprotection.BackupVaultResource{
			DppTrackedResource: armdataprotection.DppTrackedResource{
				Identity: &armdataprotection.DppIdentityDetails{
					Type: to.StringPtr("<type>"),
				},
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("val1"),
				},
			},
			Properties: &armdataprotection.BackupVault{
				StorageSettings: []*armdataprotection.StorageSetting{
					{
						Type:          armdataprotection.StorageSettingTypesLocallyRedundant.ToPtr(),
						DatastoreType: armdataprotection.StorageSettingStoreTypesVaultStore.ToPtr(),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BackupVaultResource.ID: %s\n", *res.ID)
}
