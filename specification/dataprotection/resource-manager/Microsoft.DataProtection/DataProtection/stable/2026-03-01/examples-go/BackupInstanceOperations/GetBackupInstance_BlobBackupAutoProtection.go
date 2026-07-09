package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-03-01/BackupInstanceOperations/GetBackupInstance_BlobBackupAutoProtection.json
func ExampleBackupInstancesClient_Get_getBackupInstanceWithBlobBackupAutoProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("54707983-993e-43de-8d94-074451394eda", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupInstancesClient().Get(ctx, "blobrg", "blobvault", "blobbackupinstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupInstancesClientGetResponse{
	// 	BackupInstanceResource: armdataprotection.BackupInstanceResource{
	// 		Name: to.Ptr("blobbackupinstance"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		ID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupInstances/blobbackupinstance"),
	// 		Properties: &armdataprotection.BackupInstance{
	// 			CurrentProtectionState: to.Ptr(armdataprotection.CurrentProtectionStateProtectionConfigured),
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				DatasourceType: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 				ResourceID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 				ResourceLocation: to.Ptr("centraluseuap"),
	// 				ResourceName: to.Ptr("blobstorageaccount"),
	// 				ResourceType: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 				ResourceURI: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				DatasourceType: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 				ResourceID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 				ResourceLocation: to.Ptr("centraluseuap"),
	// 				ResourceName: to.Ptr("blobstorageaccount"),
	// 				ResourceType: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 				ResourceURI: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 			},
	// 			FriendlyName: to.Ptr("blobbackupinstance"),
	// 			ObjectType: to.Ptr("BackupInstance"),
	// 			PolicyInfo: &armdataprotection.PolicyInfo{
	// 				PolicyID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupPolicies/blobpolicy"),
	// 				PolicyParameters: &armdataprotection.PolicyParameters{
	// 					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
	// 						&armdataprotection.BlobBackupDatasourceParametersForAutoProtection{
	// 							AutoProtectionSettings: &armdataprotection.BlobBackupRuleBasedAutoProtectionSettings{
	// 								Enabled: to.Ptr(true),
	// 								ObjectType: to.Ptr("BlobBackupRuleBasedAutoProtectionSettings"),
	// 								Rules: []*armdataprotection.BlobBackupAutoProtectionRule{
	// 									{
	// 										ObjectType: to.Ptr("BlobBackupAutoProtectionRule"),
	// 										Mode: to.Ptr(armdataprotection.BlobBackupRuleModeExclude),
	// 										Type: to.Ptr(armdataprotection.BlobBackupPatternTypePrefix),
	// 										Pattern: to.Ptr("temp-"),
	// 									},
	// 									{
	// 										ObjectType: to.Ptr("BlobBackupAutoProtectionRule"),
	// 										Mode: to.Ptr(armdataprotection.BlobBackupRuleModeExclude),
	// 										Type: to.Ptr(armdataprotection.BlobBackupPatternTypePrefix),
	// 										Pattern: to.Ptr("test-"),
	// 									},
	// 								},
	// 							},
	// 							ObjectType: to.Ptr("BlobBackupDatasourceParametersForAutoProtection"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 				Status: to.Ptr(armdataprotection.StatusProtectionConfigured),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
