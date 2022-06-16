package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/BackupStorageConfig_Patch.json
func ExampleBackupResourceStorageConfigsNonCRRClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupResourceStorageConfigsNonCRRClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Patch(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceConfig{
				StorageType:      to.Ptr(armrecoveryservicesbackup.StorageTypeLocallyRedundant),
				StorageTypeState: to.Ptr(armrecoveryservicesbackup.StorageTypeStateUnlocked),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
