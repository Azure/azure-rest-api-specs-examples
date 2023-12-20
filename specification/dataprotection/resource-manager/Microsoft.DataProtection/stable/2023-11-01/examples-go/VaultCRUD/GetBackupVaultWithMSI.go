package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/VaultCRUD/GetBackupVaultWithMSI.json
func ExampleBackupVaultsClient_Get_getBackupVaultWithMsi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupVaultsClient().Get(ctx, "SampleResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("c009b9a0-0024-417c-83cd-025d3776045d"),
	// 		TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		FeatureSettings: &armdataprotection.FeatureSettings{
	// 			CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 				State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
	// 		SecuritySettings: &armdataprotection.SecuritySettings{
	// 			SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 				RetentionDurationInDays: to.Ptr[float64](14),
	// 				State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
	// 			},
	// 		},
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}
