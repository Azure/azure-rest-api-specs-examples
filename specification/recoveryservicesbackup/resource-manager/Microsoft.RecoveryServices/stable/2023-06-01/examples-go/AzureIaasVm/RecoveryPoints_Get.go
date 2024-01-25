package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d402f685809d6d08be9c0b45065cadd7d78ab870/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/RecoveryPoints_Get.json
func ExampleRecoveryPointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().Get(ctx, "rshvault", "rshhtestmdvmrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall", "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall", "26083826328862", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoveryPointResource = armrecoveryservicesbackup.RecoveryPointResource{
	// 	Name: to.Ptr("26083826328862"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rshhtestmdvmrg/providers/Microsoft.RecoveryServices/vaults/rshvault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/protectedItems/VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/recoveryPoints/26083826328862"),
	// 	Properties: &armrecoveryservicesbackup.IaasVMRecoveryPoint{
	// 		ObjectType: to.Ptr("IaasVMRecoveryPoint"),
	// 		IsInstantIlrSessionActive: to.Ptr(false),
	// 		IsManagedVirtualMachine: to.Ptr(true),
	// 		IsPrivateAccessEnabledOnAnyDisk: to.Ptr(true),
	// 		IsSourceVMEncrypted: to.Ptr(false),
	// 		OriginalStorageAccountOption: to.Ptr(false),
	// 		RecoveryPointAdditionalInfo: to.Ptr(""),
	// 		RecoveryPointMoveReadinessInfo: map[string]*armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
	// 			"ArchivedRP": &armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
	// 				IsReadyForMove: to.Ptr(true),
	// 			},
	// 		},
	// 		RecoveryPointTierDetails: []*armrecoveryservicesbackup.RecoveryPointTierInformationV2{
	// 			{
	// 				Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
	// 				Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusValid),
	// 		}},
	// 		RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-22T22:32:46.608Z"); return t}()),
	// 		RecoveryPointType: to.Ptr("CrashConsistent"),
	// 		SourceVMStorageType: to.Ptr("NormalStorage"),
	// 		VirtualMachineSize: to.Ptr("Standard_D1"),
	// 		Zones: []*string{
	// 			to.Ptr("1")},
	// 		},
	// 	}
}
