package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/GetBackupVaultsInResourceGroup.json
func ExampleBackupVaultsClient_NewGetInResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.BackupVaultResourceList = armdataprotection.BackupVaultResourceList{
		// 	Value: []*armdataprotection.BackupVaultResource{
		// 		{
		// 			Name: to.Ptr("ExampleVault1"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault1"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				BcdrSecurityLevel: to.Ptr(armdataprotection.BCDRSecurityLevelGood),
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ExampleVault2"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault2"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("SystemAssigned"),
		// 				PrincipalID: to.Ptr("c009b9a0-0024-417c-83cd-025d3776045d"),
		// 				TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				BcdrSecurityLevel: to.Ptr(armdataprotection.BCDRSecurityLevelGood),
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				MonitoringSettings: &armdataprotection.MonitoringSettings{
		// 					AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
		// 						AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
