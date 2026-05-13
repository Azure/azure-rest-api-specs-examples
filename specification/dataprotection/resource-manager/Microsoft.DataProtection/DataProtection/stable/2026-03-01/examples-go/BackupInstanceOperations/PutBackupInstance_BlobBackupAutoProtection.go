package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-03-01/BackupInstanceOperations/PutBackupInstance_BlobBackupAutoProtection.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate_createBackupInstanceWithBlobBackupAutoProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("54707983-993e-43de-8d94-074451394eda", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginCreateOrUpdate(ctx, "blobrg", "blobvault", "blobstorageaccount-blobstorageaccount-2a76f8a-c176-4f7d-819e-95157e2b0071", armdataprotection.BackupInstanceResource{
		Properties: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
				ResourceLocation: to.Ptr("centraluseuap"),
				ResourceName:     to.Ptr("blobstorageaccount"),
				ResourceType:     to.Ptr("microsoft.storage/storageAccounts"),
				ResourceURI:      to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
				ResourceLocation: to.Ptr("centraluseuap"),
				ResourceName:     to.Ptr("blobstorageaccount"),
				ResourceType:     to.Ptr("microsoft.storage/storageAccounts"),
				ResourceURI:      to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
			},
			FriendlyName: to.Ptr("blobstorageaccount\\blobbackupinstance"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupPolicies/blobpolicy"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
						&armdataprotection.BlobBackupDatasourceParametersForAutoProtection{
							AutoProtectionSettings: &armdataprotection.BlobBackupRuleBasedAutoProtectionSettings{
								Enabled:    to.Ptr(true),
								ObjectType: to.Ptr("BlobBackupRuleBasedAutoProtectionSettings"),
								Rules: []*armdataprotection.BlobBackupAutoProtectionRule{
									{
										ObjectType: to.Ptr("BlobBackupAutoProtectionRule"),
										Mode:       to.Ptr(armdataprotection.BlobBackupRuleModeExclude),
										Type:       to.Ptr(armdataprotection.BlobBackupPatternTypePrefix),
										Pattern:    to.Ptr("temp-"),
									},
									{
										ObjectType: to.Ptr("BlobBackupAutoProtectionRule"),
										Mode:       to.Ptr(armdataprotection.BlobBackupRuleModeExclude),
										Type:       to.Ptr(armdataprotection.BlobBackupPatternTypePrefix),
										Pattern:    to.Ptr("test-"),
									},
								},
							},
							ObjectType: to.Ptr("BlobBackupDatasourceParametersForAutoProtection"),
						},
					},
				},
			},
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
	// res = armdataprotection.BackupInstancesClientCreateOrUpdateResponse{
	// 	BackupInstanceResource: &armdataprotection.BackupInstanceResource{
	// 		Name: to.Ptr("2a76f8a-c176-4f7d-819e-95157e2b0077"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		ID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupInstances/2a76f8a-c176-4f7d-819e-95157e2b0077"),
	// 		Properties: &armdataprotection.BackupInstance{
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				DatasourceType: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 				ResourceID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 				ResourceLocation: to.Ptr("centraluseuap"),
	// 				ResourceName: to.Ptr("blobstorageaccount"),
	// 				ResourceType: to.Ptr("microsoft.storage/storageAccounts"),
	// 				ResourceURI: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				DatasourceType: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 				ResourceID: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 				ResourceLocation: to.Ptr("centraluseuap"),
	// 				ResourceName: to.Ptr("blobstorageaccount"),
	// 				ResourceType: to.Ptr("microsoft.storage/storageAccounts"),
	// 				ResourceURI: to.Ptr("/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"),
	// 			},
	// 			FriendlyName: to.Ptr("blobstorageaccount\\blobbackupinstance"),
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
	// 				Status: to.Ptr(armdataprotection.Status("NotProtected")),
	// 			},
	// 			ProvisioningState: to.Ptr("Provisioned"),
	// 		},
	// 	},
	// }
}
