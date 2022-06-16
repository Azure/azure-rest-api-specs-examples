package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/BackupStorageConfig_Patch.json
func ExampleBackupResourceStorageConfigsNonCRRClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupResourceStorageConfigsNonCRRClient("<subscription-id>", cred, nil)
	_, err = client.Patch(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceConfig{
				StorageType:      armrecoveryservicesbackup.StorageTypeLocallyRedundant.ToPtr(),
				StorageTypeState: armrecoveryservicesbackup.StorageTypeStateUnlocked.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
