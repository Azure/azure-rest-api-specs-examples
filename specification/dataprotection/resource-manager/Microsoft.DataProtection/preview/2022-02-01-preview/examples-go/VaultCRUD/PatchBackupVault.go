package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/VaultCRUD/PatchBackupVault.json
func ExampleBackupVaultsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupVaultsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armdataprotection.PatchResourceRequestInput{
			Properties: &armdataprotection.PatchBackupVaultInput{
				MonitoringSettings: &armdataprotection.MonitoringSettings{
					AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
						AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
					},
				},
			},
			Tags: map[string]*string{
				"newKey": to.Ptr("newVal"),
			},
		},
		&armdataprotection.BackupVaultsClientBeginUpdateOptions{ResumeToken: ""})
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
