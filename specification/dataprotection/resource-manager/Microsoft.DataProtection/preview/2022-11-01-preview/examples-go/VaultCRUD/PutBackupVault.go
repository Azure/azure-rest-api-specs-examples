package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/VaultCRUD/PutBackupVault.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_createBackupVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupVaultsClient("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Identity: &armdataprotection.DppIdentityDetails{
			Type: to.Ptr("None"),
		},
		Location: to.Ptr("WestUS"),
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
