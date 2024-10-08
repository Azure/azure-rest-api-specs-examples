package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/BackupInstanceOperations/ListBackupInstancesExtensionRouting.json
func ExampleBackupInstancesExtensionRoutingClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupInstancesExtensionRoutingClient().NewListPager("subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk", nil)
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
		// page.BackupInstanceResourceList = armdataprotection.BackupInstanceResourceList{
		// 	Value: []*armdataprotection.BackupInstanceResource{
		// 		{
		// 			Name: to.Ptr("testDiskBI1-testDiskBI1-7664c12f-4d0a-440f-a0dc-b64f708b10e3"),
		// 			Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
		// 			ID: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk/providers/Microsoft.DataProtection/backupInstances/testDiskBI1-testDiskBI1-7664c12f-4d0a-440f-a0dc-b64f708b10e3"),
		// 			Properties: &armdataprotection.BackupInstance{
		// 				CurrentProtectionState: to.Ptr(armdataprotection.CurrentProtectionStateProtectionConfigured),
		// 				DataSourceInfo: &armdataprotection.Datasource{
		// 					DatasourceType: to.Ptr("Microsoft.Compute/disks"),
		// 					ObjectType: to.Ptr("Datasource"),
		// 					ResourceID: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk"),
		// 					ResourceLocation: to.Ptr("eastus2euap"),
		// 					ResourceName: to.Ptr("testDisk"),
		// 					ResourceType: to.Ptr("Microsoft.Compute/disks"),
		// 					ResourceURI: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk"),
		// 				},
		// 				DataSourceSetInfo: &armdataprotection.DatasourceSet{
		// 					DatasourceType: to.Ptr("Microsoft.Compute/disks"),
		// 					ObjectType: to.Ptr("DatasourceSet"),
		// 					ResourceID: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk"),
		// 					ResourceLocation: to.Ptr("eastus2euap"),
		// 					ResourceName: to.Ptr("testDisk"),
		// 					ResourceType: to.Ptr("Microsoft.Compute/disks"),
		// 					ResourceURI: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk"),
		// 				},
		// 				FriendlyName: to.Ptr("testDisk"),
		// 				ObjectType: to.Ptr("BackupInstance"),
		// 				PolicyInfo: &armdataprotection.PolicyInfo{
		// 					PolicyID: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/policyRG/providers/Microsoft.DataProtection/backupVaults/jeczrsecy/backupPolicies/disk"),
		// 					PolicyParameters: &armdataprotection.PolicyParameters{
		// 						DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
		// 							&armdataprotection.AzureOperationalStoreParameters{
		// 								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
		// 								ObjectType: to.Ptr("AzureOperationalStoreParameters"),
		// 								ResourceGroupID: to.Ptr("/subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/policyRG"),
		// 						}},
		// 					},
		// 					PolicyVersion: to.Ptr(""),
		// 				},
		// 				ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
		// 					Status: to.Ptr(armdataprotection.StatusProtectionConfigured),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
