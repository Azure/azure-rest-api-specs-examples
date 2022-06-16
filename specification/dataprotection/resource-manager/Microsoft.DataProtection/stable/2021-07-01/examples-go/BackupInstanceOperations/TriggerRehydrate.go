package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/TriggerRehydrate.json
func ExampleBackupInstancesClient_BeginTriggerRehydrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginTriggerRehydrate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.AzureBackupRehydrationRequest{
			RecoveryPointID:              to.StringPtr("<recovery-point-id>"),
			RehydrationPriority:          armdataprotection.RehydrationPriorityHigh.ToPtr(),
			RehydrationRetentionDuration: to.StringPtr("<rehydration-retention-duration>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
