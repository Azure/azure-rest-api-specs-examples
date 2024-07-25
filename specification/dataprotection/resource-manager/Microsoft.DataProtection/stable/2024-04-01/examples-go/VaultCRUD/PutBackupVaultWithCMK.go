package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/PutBackupVaultWithCMK.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_createBackupVaultWithCmk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Location: to.Ptr("WestUS"),
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Identity: &armdataprotection.DppIdentityDetails{
			Type: to.Ptr("None"),
		},
		Properties: &armdataprotection.BackupVault{
			MonitoringSettings: &armdataprotection.MonitoringSettings{
				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
				},
			},
			SecuritySettings: &armdataprotection.SecuritySettings{
				EncryptionSettings: &armdataprotection.EncryptionSettings{
					InfrastructureEncryption: to.Ptr(armdataprotection.InfrastructureEncryptionStateEnabled),
					KekIdentity: &armdataprotection.CmkKekIdentity{
						IdentityID:   to.Ptr("/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi"),
						IdentityType: to.Ptr(armdataprotection.IdentityTypeUserAssigned),
					},
					KeyVaultProperties: &armdataprotection.CmkKeyVaultProperties{
						KeyURI: to.Ptr("https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3"),
					},
					State: to.Ptr(armdataprotection.EncryptionStateEnabled),
				},
				ImmutabilitySettings: &armdataprotection.ImmutabilitySettings{
					State: to.Ptr(armdataprotection.ImmutabilityStateDisabled),
				},
				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
					RetentionDurationInDays: to.Ptr[float64](0),
					State:                   to.Ptr(armdataprotection.SoftDeleteStateOff),
				},
			},
			StorageSettings: []*armdataprotection.StorageSetting{
				{
					Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
				}},
		},
	}, &armdataprotection.BackupVaultsClientBeginCreateOrUpdateOptions{XMSAuthorizationAuxiliary: nil})
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
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armdataprotection.DppIdentityDetails{
	// 		Type: to.Ptr("None"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}
