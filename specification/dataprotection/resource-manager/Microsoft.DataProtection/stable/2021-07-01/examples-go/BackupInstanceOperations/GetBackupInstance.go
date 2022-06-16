package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/GetBackupInstance.json
func ExampleBackupInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-instance-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BackupInstanceResource.ID: %s\n", *res.ID)
}
