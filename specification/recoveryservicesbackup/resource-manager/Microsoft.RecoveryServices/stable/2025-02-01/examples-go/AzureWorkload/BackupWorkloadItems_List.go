package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/BackupWorkloadItems_List.json
func ExampleBackupWorkloadItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupWorkloadItemsClient().NewListPager("suchandr-seacan-rsv", "testRg", "Azure", "VMAppContainer;Compute;bvtdtestag;sqlserver-1", &armrecoveryservicesbackup.BackupWorkloadItemsClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureWorkload'"),
		SkipToken: nil,
	})
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
		// page.WorkloadItemResourceList = armrecoveryservicesbackup.WorkloadItemResourceList{
		// 	Value: []*armrecoveryservicesbackup.WorkloadItemResource{
		// 		{
		// 			Name: to.Ptr("SQLInstance;MSSQLSERVER"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/items"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/testRg/providers/Microsoft.RecoveryServices/vaults/suchandr-seacan-rsv/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;bvtdtestag;sqlserver-1/protectableItems/SQLInstance;MSSQLSERVER"),
		// 			Properties: &armrecoveryservicesbackup.AzureVMWorkloadSQLInstanceWorkloadItem{
		// 				BackupManagementType: to.Ptr("AzureWorkload"),
		// 				FriendlyName: to.Ptr("MSSQLSERVER"),
		// 				ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStatusNotProtected),
		// 				WorkloadItemType: to.Ptr("SQLInstance"),
		// 				WorkloadType: to.Ptr("SQL"),
		// 				IsAutoProtectable: to.Ptr(true),
		// 				ParentName: to.Ptr("MSSQLSERVER"),
		// 				ServerName: to.Ptr("sqlserver-1.contoso.com"),
		// 				SubWorkloadItemCount: to.Ptr[int32](3),
		// 				Subinquireditemcount: to.Ptr[int32](0),
		// 				DataDirectoryPaths: []*armrecoveryservicesbackup.SQLDataDirectory{
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.SQLDataDirectoryTypeData),
		// 						Path: to.Ptr("F:\\DATA\\"),
		// 					},
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.SQLDataDirectoryTypeLog),
		// 						Path: to.Ptr("F:\\LOG\\"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
