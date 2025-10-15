package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/BackupInstanceOperations/PutBackupInstance_ResourceGuardEnabled.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate_createBackupInstanceToPerformCriticalOperationWithMua() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginCreateOrUpdate(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", armdataprotection.BackupInstanceResource{
		Properties: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("testdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
				ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
				SecretStoreResource: &armdataprotection.SecretStoreResource{
					SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
					URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
				},
			},
			FriendlyName: to.Ptr("harshitbi2"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/Backupvaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
						&armdataprotection.AzureOperationalStoreParameters{
							DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
							ObjectType:      to.Ptr("AzureOperationalStoreParameters"),
							ResourceGroupID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest"),
						},
					},
				},
			},
			ResourceGuardOperationRequests: []*string{
				to.Ptr("/subscriptions/38304e13-357e-405e-9e9a-220351dcce8c/resourcegroups/ankurResourceGuard1/providers/Microsoft.DataProtection/resourceGuards/ResourceGuard38-1/dppModifyPolicy/default"),
			},
			ValidationType: to.Ptr(armdataprotection.ValidationTypeShallowValidation),
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
	// 		Name: to.Ptr("harshitbi2"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/backupInstances/harshitbi2"),
	// 		Properties: &armdataprotection.BackupInstance{
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				DatasourceType: to.Ptr("OssDB"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 				ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
	// 				ResourceLocation: to.Ptr(""),
	// 				ResourceName: to.Ptr("testdb"),
	// 				ResourceType: to.Ptr("OssDB"),
	// 				ResourceURI: to.Ptr(""),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				DatasourceType: to.Ptr("OssDB"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 				ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
	// 				ResourceLocation: to.Ptr(""),
	// 				ResourceName: to.Ptr("viveksipgtest"),
	// 				ResourceType: to.Ptr("OssDB"),
	// 				ResourceURI: to.Ptr(""),
	// 			},
	// 			FriendlyName: to.Ptr("harshitbi2"),
	// 			ObjectType: to.Ptr("BackupInstance"),
	// 			PolicyInfo: &armdataprotection.PolicyInfo{
	// 				PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
	// 				PolicyParameters: &armdataprotection.PolicyParameters{
	// 					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
	// 						&armdataprotection.AzureOperationalStoreParameters{
	// 							DataStoreType: to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
	// 							ObjectType: to.Ptr("AzureOperationalStoreParameters"),
	// 							ResourceGroupID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest"),
	// 						},
	// 					},
	// 				},
	// 				PolicyVersion: to.Ptr("3.2"),
	// 			},
	// 			ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 				Status: to.Ptr(armdataprotection.Status("NotProtected")),
	// 			},
	// 			ProvisioningState: to.Ptr("Provisioned"),
	// 		},
	// 	},
	// }
}
