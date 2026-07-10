package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-03-01/DeletedBackupVaults_Get.json
func ExampleDeletedBackupVaultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedBackupVaultsClient().Get(ctx, "westus", "deleted-vault-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.DeletedBackupVaultsClientGetResponse{
	// 	DeletedBackupVaultResource: armdataprotection.DeletedBackupVaultResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DataProtection/locations/westus/deletedBackupVaults/deleted-vault-01"),
	// 		Name: to.Ptr("deleted-vault-01"),
	// 		Type: to.Ptr("Microsoft.DataProtection/deletedBackupVaults"),
	// 		Properties: &armdataprotection.DeletedBackupVault{
	// 			OriginalBackupVaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DataProtection/backupVaults/sample-vault"),
	// 			OriginalBackupVaultName: to.Ptr("sample-vault"),
	// 			OriginalBackupVaultResourcePath: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DataProtection/backupVaults/sample-vault"),
	// 			ResourceDeletionInfo: &armdataprotection.ResourceDeletionInfo{
	// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-18T10:00:00.000Z"); return t}()),
	// 				ScheduledPurgeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-18T10:00:00.000Z"); return t}()),
	// 				DeleteActivityID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 			ResourceMoveState: to.Ptr(armdataprotection.ResourceMoveStateUnknown),
	// 			SecuritySettings: &armdataprotection.SecuritySettings{
	// 				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 					State: to.Ptr(armdataprotection.SoftDeleteStateOn),
	// 					RetentionDurationInDays: to.Ptr[float64](14),
	// 				},
	// 				ImmutabilitySettings: &armdataprotection.ImmutabilitySettings{
	// 					State: to.Ptr(armdataprotection.ImmutabilityStateDisabled),
	// 				},
	// 			},
	// 			StorageSettings: []*armdataprotection.StorageSetting{
	// 				{
	// 					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 					Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				},
	// 			},
	// 			IsVaultProtectedByResourceGuard: to.Ptr(false),
	// 			FeatureSettings: &armdataprotection.FeatureSettings{
	// 				CrossSubscriptionRestoreSettings: &armdataprotection.CrossSubscriptionRestoreSettings{
	// 					State: to.Ptr(armdataprotection.CrossSubscriptionRestoreStateDisabled),
	// 				},
	// 				CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 					State: to.Ptr(armdataprotection.CrossRegionRestoreStateDisabled),
	// 				},
	// 			},
	// 			SecureScore: to.Ptr(armdataprotection.SecureScoreLevelMaximum),
	// 			ResourceGuardOperationRequests: []*string{
	// 			},
	// 			ReplicatedRegions: []*string{
	// 			},
	// 		},
	// 	},
	// }
}
