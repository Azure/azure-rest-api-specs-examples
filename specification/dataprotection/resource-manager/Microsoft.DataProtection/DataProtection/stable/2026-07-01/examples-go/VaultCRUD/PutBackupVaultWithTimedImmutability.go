package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-07-01/VaultCRUD/PutBackupVaultWithTimedImmutability.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_createBackupVaultWithTimedImmutability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Location: to.Ptr("WestUS"),
		Properties: &armdataprotection.BackupVault{
			SecuritySettings: &armdataprotection.SecuritySettings{
				ImmutabilitySettings: &armdataprotection.ImmutabilitySettings{
					State: to.Ptr(armdataprotection.ImmutabilityStateUnlocked),
					Configuration: &armdataprotection.ImmutabilityConfiguration{
						Type:           to.Ptr(armdataprotection.ImmutabilityTypeTimeBased),
						DurationInDays: to.Ptr[int32](30),
					},
				},
				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
					RetentionDurationInDays: to.Ptr[float64](14),
					State:                   to.Ptr(armdataprotection.SoftDeleteStateOn),
				},
			},
			StorageSettings: []*armdataprotection.StorageSetting{
				{
					Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
				},
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Identity: &armdataprotection.DppIdentityDetails{
			Type: to.Ptr("None"),
		},
	}, nil)
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
	// 		ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/swaggerExample"),
	// 		Name: to.Ptr("swaggerExample"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults"),
	// 		Location: to.Ptr("WestUS"),
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 		},
	// 		Identity: &armdataprotection.DppIdentityDetails{
	// 			Type: to.Ptr("None"),
	// 		},
	// 		Properties: &armdataprotection.BackupVault{
	// 			ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 			SecuritySettings: &armdataprotection.SecuritySettings{
	// 				ImmutabilitySettings: &armdataprotection.ImmutabilitySettings{
	// 					State: to.Ptr(armdataprotection.ImmutabilityStateUnlocked),
	// 					Configuration: &armdataprotection.ImmutabilityConfiguration{
	// 						Type: to.Ptr(armdataprotection.ImmutabilityTypeTimeBased),
	// 						DurationInDays: to.Ptr[int32](30),
	// 					},
	// 				},
	// 				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 					RetentionDurationInDays: to.Ptr[float64](14),
	// 					State: to.Ptr(armdataprotection.SoftDeleteStateOn),
	// 				},
	// 			},
	// 			StorageSettings: []*armdataprotection.StorageSetting{
	// 				{
	// 					Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
