package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/VaultCRUD/GetBackupVaultsInResourceGroup.json
func ExampleBackupVaultsClient_NewGetInResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupVaultsClient().NewGetInResourceGroupPager("SampleResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdataprotection.BackupVaultsClientGetInResourceGroupResponse{
		// 	BackupVaultResourceList: armdataprotection.BackupVaultResourceList{
		// 		Value: []*armdataprotection.BackupVaultResource{
		// 			{
		// 				Name: to.Ptr("ExampleVault1"),
		// 				Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 				ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault1"),
		// 				Location: to.Ptr("WestUS"),
		// 				Properties: &armdataprotection.BackupVault{
		// 					BcdrSecurityLevel: to.Ptr(armdataprotection.BCDRSecurityLevelGood),
		// 					FeatureSettings: &armdataprotection.FeatureSettings{
		// 						CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 							State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 					SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 					SecuritySettings: &armdataprotection.SecuritySettings{
		// 						SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 							RetentionDurationInDays: to.Ptr[float64](14),
		// 							State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 						},
		// 					},
		// 					StorageSettings: []*armdataprotection.StorageSetting{
		// 						{
		// 							Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 							DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("val1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ExampleVault2"),
		// 				Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 				ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault2"),
		// 				Location: to.Ptr("WestUS"),
		// 				Properties: &armdataprotection.BackupVault{
		// 					BcdrSecurityLevel: to.Ptr(armdataprotection.BCDRSecurityLevelGood),
		// 					FeatureSettings: &armdataprotection.FeatureSettings{
		// 						CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 							State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 						},
		// 					},
		// 					MonitoringSettings: &armdataprotection.MonitoringSettings{
		// 						AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
		// 							AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 					SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 					SecuritySettings: &armdataprotection.SecuritySettings{
		// 						SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 							RetentionDurationInDays: to.Ptr[float64](14),
		// 							State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 						},
		// 					},
		// 					StorageSettings: []*armdataprotection.StorageSetting{
		// 						{
		// 							Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 							DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("val1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
