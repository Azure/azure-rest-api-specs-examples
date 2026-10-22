package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-10-01/AzureIaasVm/GetRPExtendedInfoOperationResult.json
func ExampleRecoveryPointsClient_GetRPExtendedInfoOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().GetRPExtendedInfoOperationResult(ctx, "testRG", "testVault", "Azure", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.RecoveryPointsClientGetRPExtendedInfoOperationResultResponse{
	// 	RPExtendedInfoResourceList: armrecoveryservicesbackup.RPExtendedInfoResourceList{
	// 		Value: []*armrecoveryservicesbackup.RPExtendedInfoResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;testRG;testvmName/protectedItems/VM;iaasvmcontainerv2;testRG;testvmName/recoveryPoints/348916168024334"),
	// 				Name: to.Ptr("348916168024334"),
	// 				Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
	// 				Properties: &armrecoveryservicesbackup.IaasVMRPExtendedInfo{
	// 					ObjectType: to.Ptr("IaasVMRPExtendedInfo"),
	// 					OSDisk: &armrecoveryservicesbackup.DiskInformation{
	// 						DiskSizeInGB: to.Ptr[int32](127),
	// 						Lun: to.Ptr[int32](-1),
	// 						Name: to.Ptr("testvmName_OsDisk_1"),
	// 						StorageType: to.Ptr("Premium_LRS"),
	// 					},
	// 					DataDisks: []*armrecoveryservicesbackup.DiskInformation{
	// 						{
	// 							DiskSizeInGB: to.Ptr[int32](1024),
	// 							Lun: to.Ptr[int32](0),
	// 							Name: to.Ptr("testvmName_DataDisk_0"),
	// 							StorageType: to.Ptr("StandardSSD_LRS"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
