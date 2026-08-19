package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-07-01/AzureIaasVm/RecoveryPoints_Get.json
func ExampleRecoveryPointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
	// res = armrecoveryservicesbackup.RecoveryPointsClientGetResponse{
	// 	RecoveryPointResource: armrecoveryservicesbackup.RecoveryPointResource{
	// 		Name: to.Ptr("26083826328862"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rshhtestmdvmrg/providers/Microsoft.RecoveryServices/vaults/rshvault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/protectedItems/VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall/recoveryPoints/26083826328862"),
	// 		Properties: &armrecoveryservicesbackup.IaasVMRecoveryPoint{
	// 			DataDiskMetadata: &armrecoveryservicesbackup.DataDiskDetails{
	// 				EncryptedDataDisks: []*armrecoveryservicesbackup.DiskDetails{
	// 					{
	// 						Lun: to.Ptr[int32](0),
	// 						DiskName: to.Ptr("cvm-06170038-DataDisk-0"),
	// 					},
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						DiskName: to.Ptr("cvm-06170038-DataDisk-1"),
	// 					},
	// 				},
	// 			},
	// 			IsInstantIlrSessionActive: to.Ptr(false),
	// 			IsManagedVirtualMachine: to.Ptr(true),
	// 			IsPrivateAccessEnabledOnAnyDisk: to.Ptr(true),
	// 			IsSourceVMEncrypted: to.Ptr(false),
	// 			ObjectType: to.Ptr("IaasVMRecoveryPoint"),
	// 			OriginalStorageAccountOption: to.Ptr(false),
	// 			RecoveryPointAdditionalInfo: to.Ptr(""),
	// 			RecoveryPointMoveReadinessInfo: map[string]*armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
	// 				"ArchivedRP": &armrecoveryservicesbackup.RecoveryPointMoveReadinessInfo{
	// 					IsReadyForMove: to.Ptr(true),
	// 				},
	// 			},
	// 			RecoveryPointTierDetails: []*armrecoveryservicesbackup.RecoveryPointTierInformationV2{
	// 				{
	// 					Type: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
	// 					Status: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierStatusValid),
	// 				},
	// 			},
	// 			RecoveryPointTime: to.Ptr(time.Date(2017, time.November, 22, 22, 32, 46, 608847200, time.UTC)),
	// 			RecoveryPointType: to.Ptr("CrashConsistent"),
	// 			SourceVMStorageType: to.Ptr("NormalStorage"),
	// 			VirtualMachineSize: to.Ptr("Standard_D1"),
	// 			RecoveryPointProperties: &armrecoveryservicesbackup.RecoveryPointProperties{
	// 				ExpiryTime: to.Ptr("2020-11-22T22:32:46.6088472Z"),
	// 				RuleName: to.Ptr("DefaultRule"),
	// 				ImmutabilityProperties: &armrecoveryservicesbackup.RecoveryPointImmutabilityProperties{
	// 					IsImmutable: to.Ptr(true),
	// 					ExpiryTime: to.Ptr(time.Date(2020, time.November, 22, 22, 32, 46, 608847200, time.UTC)),
	// 				},
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1"),
	// 			},
	// 			ThreatStatus: to.Ptr(armrecoveryservicesbackup.ThreatStatusHealthy),
	// 			ThreatInfo: []*armrecoveryservicesbackup.ThreatInfo{
	// 				{
	// 					ThreatTitle: to.Ptr("MDC threat title"),
	// 					ThreatDescription: to.Ptr("Threat description"),
	// 					LastUpdatedTime: to.Ptr(time.Date(2025, time.January, 22, 22, 32, 46, 608847200, time.UTC)),
	// 					ThreatState: to.Ptr(armrecoveryservicesbackup.ThreatStateActive),
	// 					ThreatStartTime: to.Ptr(time.Date(2024, time.November, 22, 22, 32, 46, 608847200, time.UTC)),
	// 					ThreatEndTime: to.Ptr(time.Date(2024, time.November, 23, 17, 13, 23, 604547200, time.UTC)),
	// 					ThreatURI: to.Ptr("https://portal.azure.com/#blade/Microsoft_Azure_Security_AzureDefenderForData/AlertBlade/alertId/00000000-0000-0000-0000-000000000000/subscriptionId/00000000-0000-0000-0000-000000000000/resourceGroup/Sample-RG/referencedFrom/alertDeepLink/location/centralus"),
	// 					ThreatSeverity: to.Ptr(armrecoveryservicesbackup.ThreatSeverityInformational),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
