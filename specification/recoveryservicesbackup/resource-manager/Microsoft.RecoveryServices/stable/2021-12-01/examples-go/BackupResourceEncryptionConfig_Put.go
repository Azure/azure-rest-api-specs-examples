package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/BackupResourceEncryptionConfig_Put.json
func ExampleBackupResourceEncryptionConfigsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewBackupResourceEncryptionConfigsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Update(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceEncryptionConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceEncryptionConfig{
				EncryptionAtRestType:          to.Ptr(armrecoveryservicesbackup.EncryptionAtRestTypeCustomerManaged),
				InfrastructureEncryptionState: to.Ptr(armrecoveryservicesbackup.InfrastructureEncryptionState("true")),
				KeyURI:                        to.Ptr("<key-uri>"),
				SubscriptionID:                to.Ptr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
