package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-03-01/PutBackupVaultWithUndelete.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_restoreASoftDeletedBackupVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Location: to.Ptr("WestUS"),
		Properties: &armdataprotection.BackupVault{
			MonitoringSettings: &armdataprotection.MonitoringSettings{
				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
				},
			},
			SecuritySettings: &armdataprotection.SecuritySettings{
				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
					RetentionDurationInDays: to.Ptr[float64](14),
					State:                   to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
				},
				ImmutabilitySettings: &armdataprotection.ImmutabilitySettings{
					State: to.Ptr(armdataprotection.ImmutabilityStateDisabled),
				},
			},
			StorageSettings: []*armdataprotection.StorageSetting{
				{
					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
					Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
				},
			},
			FeatureSettings: &armdataprotection.FeatureSettings{
				CrossSubscriptionRestoreSettings: &armdataprotection.CrossSubscriptionRestoreSettings{
					State: to.Ptr(armdataprotection.CrossSubscriptionRestoreStateDisabled),
				},
				CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
					State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
				},
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
	}, &armdataprotection.BackupVaultsClientBeginCreateOrUpdateOptions{
		XMSDeletedVaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DataProtection/locations/WestUS/deletedVaults/swaggerExample")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupVaultsClientCreateOrUpdateResponse{
	// 	BackupVaultResource: armdataprotection.BackupVaultResource{
	// 		Location: to.Ptr("WestUS"),
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/swaggerExample"),
	// 		Name: to.Ptr("swaggerExample"),
	// 		Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
	// 		Properties: &armdataprotection.BackupVault{
	// 			ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 			ResourceMoveState: to.Ptr(armdataprotection.ResourceMoveStateUnknown),
	// 			MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 				},
	// 			},
	// 			SecuritySettings: &armdataprotection.SecuritySettings{
	// 				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 					RetentionDurationInDays: to.Ptr[float64](14),
	// 					State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
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
	// 			FeatureSettings: &armdataprotection.FeatureSettings{
	// 				CrossSubscriptionRestoreSettings: &armdataprotection.CrossSubscriptionRestoreSettings{
	// 					State: to.Ptr(armdataprotection.CrossSubscriptionRestoreStateDisabled),
	// 				},
	// 				CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 					State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 				},
	// 			},
	// 			SecureScore: to.Ptr(armdataprotection.SecureScoreLevelMaximum),
	// 			IsVaultProtectedByResourceGuard: to.Ptr(false),
	// 			ResourceGuardOperationRequests: []*string{
	// 			},
	// 			ReplicatedRegions: []*string{
	// 			},
	// 		},
	// 	},
	// }
}
