package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/RecoveryPointsRecommendedForMove_List.json
func ExampleRecoveryPointsRecommendedForMoveClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecoveryPointsRecommendedForMoveClient().NewListPager("rshvault", "rshhtestmdvmrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall", "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall", armrecoveryservicesbackup.ListRecoveryPointsRecommendedForMoveRequest{
		ExcludedRPList: []*string{
			to.Ptr("348916168024334"),
			to.Ptr("348916168024335")},
		ObjectType: to.Ptr("ListRecoveryPointsRecommendedForMoveRequest"),
	}, nil)
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
		// page.RecoveryPointResourceList = armrecoveryservicesbackup.RecoveryPointResourceList{
		// 	Value: []*armrecoveryservicesbackup.RecoveryPointResource{
		// 		{
		// 			Name: to.Ptr("22244821112382"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rshhtestmdvmrg/providers/Microsoft.RecoveryServices/vaults/rshvault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/protectedItems/VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/recoveryPoints/22244821112382"),
		// 			Properties: &armrecoveryservicesbackup.IaasVMRecoveryPoint{
		// 				ObjectType: to.Ptr("IaasVMRecoveryPoint"),
		// 				IsInstantIlrSessionActive: to.Ptr(false),
		// 				IsManagedVirtualMachine: to.Ptr(true),
		// 				IsSourceVMEncrypted: to.Ptr(false),
		// 				OriginalStorageAccountOption: to.Ptr(false),
		// 				RecoveryPointAdditionalInfo: to.Ptr(""),
		// 				RecoveryPointMoveReadinessInfo: map[string]*armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
		// 					"ArchivedRP": &armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
		// 						IsReadyForMove: to.Ptr(true),
		// 					},
		// 				},
		// 				RecoveryPointTierDetails: []*armrecoveryservicesbackup.RecoveryPointTierInformationV2{
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeInstantRP),
		// 						Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusDeleted),
		// 					},
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
		// 						Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusValid),
		// 				}},
		// 				RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-21T22:48:25.435Z"); return t}()),
		// 				RecoveryPointType: to.Ptr("CrashConsistent"),
		// 				SourceVMStorageType: to.Ptr("NormalStorage"),
		// 				VirtualMachineSize: to.Ptr("Standard_D1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("24977149827250"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rshhtestmdvmrg/providers/Microsoft.RecoveryServices/vaults/rshvault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/protectedItems/VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/recoveryPoints/24977149827250"),
		// 			Properties: &armrecoveryservicesbackup.IaasVMRecoveryPoint{
		// 				ObjectType: to.Ptr("IaasVMRecoveryPoint"),
		// 				IsInstantIlrSessionActive: to.Ptr(false),
		// 				IsManagedVirtualMachine: to.Ptr(true),
		// 				IsSourceVMEncrypted: to.Ptr(false),
		// 				OriginalStorageAccountOption: to.Ptr(false),
		// 				RecoveryPointAdditionalInfo: to.Ptr(""),
		// 				RecoveryPointMoveReadinessInfo: map[string]*armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
		// 					"ArchivedRP": &armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
		// 						IsReadyForMove: to.Ptr(true),
		// 					},
		// 				},
		// 				RecoveryPointTierDetails: []*armrecoveryservicesbackup.RecoveryPointTierInformationV2{
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeInstantRP),
		// 						Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusDeleted),
		// 					},
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
		// 						Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusDeleted),
		// 					},
		// 					{
		// 						Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
		// 						ExtendedInfo: map[string]*string{
		// 							"RehydratedRPExpiryTime": to.Ptr("2020-12-21T22:48:25.4353958Z"),
		// 						},
		// 						Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusRehydrated),
		// 				}},
		// 				RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-20T22:49:44.331Z"); return t}()),
		// 				RecoveryPointType: to.Ptr("CrashConsistent"),
		// 				SourceVMStorageType: to.Ptr("NormalStorage"),
		// 				VirtualMachineSize: to.Ptr("Standard_D1"),
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 				},
		// 		}},
		// 	}
	}
}
