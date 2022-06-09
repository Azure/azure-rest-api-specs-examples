```go
package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/VaultCRUD/PutBackupVault.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate() {
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
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armdataprotection.BackupVaultResource{
			Identity: &armdataprotection.DppIdentityDetails{
				Type: to.Ptr("<type>"),
			},
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			Properties: &armdataprotection.BackupVault{
				MonitoringSettings: &armdataprotection.MonitoringSettings{
					AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
						AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
					},
				},
				StorageSettings: []*armdataprotection.StorageSetting{
					{
						Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
					}},
			},
		},
		&armdataprotection.BackupVaultsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataprotection%2Farmdataprotection%2Fv0.4.0/sdk/resourcemanager/dataprotection/armdataprotection/README.md) on how to add the SDK to your project and authenticate.
