package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/BackupProtectionContainers_UsageSummary_Get.json
func ExampleBackupUsageSummariesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupUsageSummariesClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<vault-name>",
		"<resource-group-name>",
		&armrecoveryservicesbackup.BackupUsageSummariesListOptions{Filter: to.StringPtr("<filter>"),
			SkipToken: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
}
