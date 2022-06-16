package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureWorkload/BackupProtectionIntent_Delete.json
func ExampleProtectionIntentClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionIntentClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<intent-object-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
