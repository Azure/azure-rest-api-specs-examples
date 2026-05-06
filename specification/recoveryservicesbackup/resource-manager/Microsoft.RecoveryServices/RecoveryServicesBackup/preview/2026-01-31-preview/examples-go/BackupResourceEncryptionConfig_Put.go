package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/BackupResourceEncryptionConfig_Put.json
func ExampleBackupResourceEncryptionConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupResourceEncryptionConfigsClient().Update(ctx, "source-rsv", "test-rg", armrecoveryservicesbackup.BackupResourceEncryptionConfigResource{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.BackupResourceEncryptionConfigsClientUpdateResponse{
	// }
}
