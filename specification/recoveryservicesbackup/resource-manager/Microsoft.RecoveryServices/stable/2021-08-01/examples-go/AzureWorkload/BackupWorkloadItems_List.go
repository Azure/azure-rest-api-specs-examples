package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureWorkload/BackupWorkloadItems_List.json
func ExampleBackupWorkloadItemsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupWorkloadItemsClient("<subscription-id>", cred, nil)
	pager := client.List("<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		&armrecoveryservicesbackup.BackupWorkloadItemsListOptions{Filter: to.StringPtr("<filter>"),
			SkipToken: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadItemResource.ID: %s\n", *v.ID)
		}
	}
}
