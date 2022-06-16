package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/BackupResourceVaultConfigs_Patch.json
func ExampleBackupResourceVaultConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupResourceVaultConfigsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceVaultConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceVaultConfig{
				EnhancedSecurityState: armrecoveryservicesbackup.EnhancedSecurityStateEnabled.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BackupResourceVaultConfigResource.ID: %s\n", *res.ID)
}
