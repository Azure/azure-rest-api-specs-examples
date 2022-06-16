package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/BackupResourceEncryptionConfig_Put.json
func ExampleBackupResourceEncryptionConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupResourceEncryptionConfigsClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceEncryptionConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceEncryptionConfig{
				EncryptionAtRestType:          armrecoveryservicesbackup.EncryptionAtRestTypeCustomerManaged.ToPtr(),
				InfrastructureEncryptionState: armrecoveryservicesbackup.InfrastructureEncryptionStateInvalid.ToPtr(),
				KeyURI:                        to.StringPtr("<key-uri>"),
				SubscriptionID:                to.StringPtr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
