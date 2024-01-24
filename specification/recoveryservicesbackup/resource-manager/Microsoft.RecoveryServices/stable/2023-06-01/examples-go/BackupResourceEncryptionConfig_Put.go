package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d402f685809d6d08be9c0b45065cadd7d78ab870/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/BackupResourceEncryptionConfig_Put.json
func ExampleBackupResourceEncryptionConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackupResourceEncryptionConfigsClient().Update(ctx, "source-rsv", "test-rg", armrecoveryservicesbackup.BackupResourceEncryptionConfigResource{
		Properties: &armrecoveryservicesbackup.BackupResourceEncryptionConfig{
			EncryptionAtRestType:          to.Ptr(armrecoveryservicesbackup.EncryptionAtRestTypeCustomerManaged),
			InfrastructureEncryptionState: to.Ptr(armrecoveryservicesbackup.InfrastructureEncryptionState("true")),
			KeyURI:                        to.Ptr("https://gktestkv1.vault.azure.net/keys/Test1/ed2e8cdc7f86477ebf0c6462b504a9ed"),
			SubscriptionID:                to.Ptr("1a2311d9-66f5-47d3-a9fb-7a37da63934b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
