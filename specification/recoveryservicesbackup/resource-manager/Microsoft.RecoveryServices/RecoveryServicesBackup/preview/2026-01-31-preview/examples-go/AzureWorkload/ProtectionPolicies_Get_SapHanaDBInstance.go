package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/AzureWorkload/ProtectionPolicies_Get_SapHanaDBInstance.json
func ExampleProtectionPoliciesClient_Get_getSapHanaDbInstanceWorkloadProtectionPolicyDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionPoliciesClient().Get(ctx, "HanaTestRsVault", "SwaggerTestRg", "testHanaSnapshotV2Policy1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.ProtectionPoliciesClientGetResponse{
	// 	ProtectionPolicyResource: armrecoveryservicesbackup.ProtectionPolicyResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/HanaTestRsVault/backupPolicies/testHanaSnapshotV2Policy1"),
	// 		Name: to.Ptr("testHanaSnapshotV2Policy1"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
	// 		Properties: &armrecoveryservicesbackup.AzureVMWorkloadProtectionPolicy{
	// 			BackupManagementType: to.Ptr("AzureWorkload"),
	// 			WorkLoadType: to.Ptr(armrecoveryservicesbackup.WorkloadTypeSAPHanaDBInstance),
	// 			VMWorkloadPolicyType: to.Ptr(armrecoveryservicesbackup.VMWorkloadPolicyTypeSnapshotV2),
	// 			Settings: &armrecoveryservicesbackup.Settings{
	// 				TimeZone: to.Ptr("UTC"),
	// 				Issqlcompression: to.Ptr(false),
	// 			},
	// 			SubProtectionPolicy: []*armrecoveryservicesbackup.SubProtectionPolicy{
	// 				{
	// 					PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeSnapshotFull),
	// 					SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
	// 						SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
	// 						ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
	// 						ScheduleRunTimes: []*time.Time{
	// 							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T03:30:00Z"); return t}()),
	// 						},
	// 						ScheduleWeeklyFrequency: to.Ptr[int32](0),
	// 					},
	// 					RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
	// 						RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
	// 						DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
	// 							RetentionTimes: []*time.Time{
	// 								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-19T20:00:00Z"); return t}()),
	// 							},
	// 							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 								Count: to.Ptr[int32](30),
	// 								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
	// 							},
	// 						},
	// 						WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
	// 							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 								to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
	// 							},
	// 							RetentionTimes: []*time.Time{
	// 								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-19T20:00:00Z"); return t}()),
	// 							},
	// 							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 								Count: to.Ptr[int32](10),
	// 								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
	// 							},
	// 						},
	// 						MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
	// 							RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 							RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 								DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 									to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
	// 								},
	// 								WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 									to.Ptr(armrecoveryservicesbackup.WeekOfMonthSecond),
	// 								},
	// 							},
	// 							RetentionTimes: []*time.Time{
	// 								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-15T20:00:00Z"); return t}()),
	// 							},
	// 							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 								Count: to.Ptr[int32](6),
	// 								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
	// 							},
	// 						},
	// 						YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
	// 							RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 							MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
	// 								to.Ptr(armrecoveryservicesbackup.MonthOfYearJanuary),
	// 							},
	// 							RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 								DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 									to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
	// 								},
	// 								WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 									to.Ptr(armrecoveryservicesbackup.WeekOfMonthLast),
	// 								},
	// 							},
	// 							RetentionTimes: []*time.Time{
	// 								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-19T20:00:00Z"); return t}()),
	// 							},
	// 							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 								Count: to.Ptr[int32](2),
	// 								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
	// 							},
	// 						},
	// 					},
	// 					SnapshotBackupAdditionalDetails: &armrecoveryservicesbackup.SnapshotBackupAdditionalDetails{
	// 						InstantRpRetentionRangeInDays: to.Ptr[int32](5),
	// 						InstantRPDetails: to.Ptr("SwaggerTestRG"),
	// 						UserAssignedManagedIdentityDetails: &armrecoveryservicesbackup.UserAssignedManagedIdentityDetails{
	// 							IdentityArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerMsiRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/SwaggerUMI"),
	// 							IdentityName: to.Ptr("SwaggerUMI"),
	// 							UserAssignedIdentityProperties: &armrecoveryservicesbackup.UserAssignedIdentityProperties{
	// 								ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 								PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProtectedItemsCount: to.Ptr[int32](0),
	// 		},
	// 	},
	// }
}
