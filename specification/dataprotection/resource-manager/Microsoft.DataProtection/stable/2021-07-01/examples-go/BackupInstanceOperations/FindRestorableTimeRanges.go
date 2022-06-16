package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/FindRestorableTimeRanges.json
func ExampleRestorableTimeRangesClient_Find() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewRestorableTimeRangesClient("<subscription-id>", cred, nil)
	res, err := client.Find(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-instance-name>",
		armdataprotection.AzureBackupFindRestorableTimeRangesRequest{
			EndTime:             to.StringPtr("<end-time>"),
			SourceDataStoreType: armdataprotection.RestoreSourceDataStoreTypeOperationalStore.ToPtr(),
			StartTime:           to.StringPtr("<start-time>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AzureBackupFindRestorableTimeRangesResponseResource.ID: %s\n", *res.ID)
}
