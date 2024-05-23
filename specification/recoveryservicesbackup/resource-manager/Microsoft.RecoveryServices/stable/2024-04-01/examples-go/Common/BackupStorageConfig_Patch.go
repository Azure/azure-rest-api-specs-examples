package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/Common/BackupStorageConfig_Patch.json
func ExampleBackupResourceStorageConfigsNonCRRClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackupResourceStorageConfigsNonCRRClient().Patch(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", armrecoveryservicesbackup.BackupResourceConfigResource{
		Properties: &armrecoveryservicesbackup.BackupResourceConfig{
			StorageType:      to.Ptr(armrecoveryservicesbackup.StorageTypeLocallyRedundant),
			StorageTypeState: to.Ptr(armrecoveryservicesbackup.StorageTypeStateUnlocked),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
