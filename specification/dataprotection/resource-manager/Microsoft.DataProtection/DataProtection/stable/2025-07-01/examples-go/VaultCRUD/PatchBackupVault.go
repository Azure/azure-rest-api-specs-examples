package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/VaultCRUD/PatchBackupVault.json
func ExampleBackupVaultsClient_BeginUpdate_patchBackupVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.PatchResourceRequestInput{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupVaultsClientUpdateResponse{
	// 	BackupVaultResource: &armdataprotection.BackupVaultResource{
	// 		Name: to.Ptr("swaggerExample"),
	// 		Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 		ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 		Location: to.Ptr("WestUS"),
	// 		Properties: &armdataprotection.BackupVault{
	// 			MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 			StorageSettings: []*armdataprotection.StorageSetting{
	// 				{
	// 					Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"newKey": to.Ptr("newVal"),
	// 		},
	// 	},
	// }
}
