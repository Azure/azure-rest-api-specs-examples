package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/BackupInstanceOperations/GetBackupInstanceOperationResult.json
func ExampleBackupInstancesClient_GetBackupInstanceOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupInstancesClient().GetBackupInstanceOperationResult(ctx, "SampleResourceGroup", "swaggerExample", "testInstance1", "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupInstanceResource = armdataprotection.BackupInstanceResource{
	// 	Name: to.Ptr("testInstance1"),
	// 	Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/swaggerExample/backupInstances/testInstance1"),
	// 	Properties: &armdataprotection.BackupInstance{
	// 		DataSourceInfo: &armdataprotection.Datasource{
	// 			DatasourceType: to.Ptr("OssDB"),
	// 			ObjectType: to.Ptr("Datasource"),
	// 			ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
	// 			ResourceLocation: to.Ptr(""),
	// 			ResourceName: to.Ptr("testdb"),
	// 			ResourceType: to.Ptr("OssDB"),
	// 			ResourceURI: to.Ptr(""),
	// 		},
	// 		DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 			DatasourceType: to.Ptr("OssDB"),
	// 			ObjectType: to.Ptr("DatasourceSet"),
	// 			ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
	// 			ResourceLocation: to.Ptr(""),
	// 			ResourceName: to.Ptr("viveksipgtest"),
	// 			ResourceType: to.Ptr("OssDB"),
	// 			ResourceURI: to.Ptr(""),
	// 		},
	// 		FriendlyName: to.Ptr("testInstance1"),
	// 		ObjectType: to.Ptr("BackupInstance"),
	// 		PolicyInfo: &armdataprotection.PolicyInfo{
	// 			PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/swaggerExample/backupPolicies/PratikPolicy1"),
	// 			PolicyVersion: to.Ptr("3.2"),
	// 		},
	// 		ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 			Status: to.Ptr(armdataprotection.StatusConfiguringProtection),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
