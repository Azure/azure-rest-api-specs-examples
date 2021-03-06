package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/TriggerBackup.json
func ExampleBackupInstancesClient_BeginAdhocBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginAdhocBackup(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.TriggerBackupRequest{
			BackupRuleOptions: &armdataprotection.AdHocBackupRuleOptions{
				RuleName: to.Ptr("<rule-name>"),
				TriggerOption: &armdataprotection.AdhocBackupTriggerOption{
					RetentionTagOverride: to.Ptr("<retention-tag-override>"),
				},
			},
		},
		&armdataprotection.BackupInstancesClientBeginAdhocBackupOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
