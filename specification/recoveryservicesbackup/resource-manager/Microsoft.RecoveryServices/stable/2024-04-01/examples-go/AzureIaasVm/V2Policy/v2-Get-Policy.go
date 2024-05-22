package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureIaasVm/V2Policy/v2-Get-Policy.json
func ExampleProtectionPoliciesClient_Get_getAzureIaasVmEnhancedProtectionPolicyDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionPoliciesClient().Get(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "v2-daily-sample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionPolicyResource = armrecoveryservicesbackup.ProtectionPolicyResource{
	// 	Name: to.Ptr("v2-daily-sample"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/v2-daily-sample"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
	// 		BackupManagementType: to.Ptr("AzureIaasVM"),
	// 		ProtectedItemsCount: to.Ptr[int32](0),
	// 		InstantRpRetentionRangeInDays: to.Ptr[int32](30),
	// 		PolicyType: to.Ptr(armrecoveryservicesbackup.IAASVMPolicyTypeV2),
	// 		RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
	// 			RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
	// 			DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
	// 				RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 					Count: to.Ptr[int32](1),
	// 					DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
	// 				},
	// 				RetentionTimes: []*time.Time{
	// 					to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T02:00:00.000Z"); return t}())},
	// 				},
	// 			},
	// 			SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicyV2{
	// 				SchedulePolicyType: to.Ptr("SimpleSchedulePolicyV2"),
	// 				DailySchedule: &armrecoveryservicesbackup.DailySchedule{
	// 					ScheduleRunTimes: []*time.Time{
	// 						to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 					},
	// 					ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
	// 				},
	// 				SnapshotConsistencyType: to.Ptr(armrecoveryservicesbackup.IaasVMSnapshotConsistencyTypeOnlyCrashConsistent),
	// 				TimeZone: to.Ptr("Pacific Standard Time"),
	// 			},
	// 		}
}
