package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d402f685809d6d08be9c0b45065cadd7d78ab870/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Hardened.json
func ExampleProtectionPoliciesClient_CreateOrUpdate_createOrUpdateAzureStorageVaultStandardProtectionPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionPoliciesClient().CreateOrUpdate(ctx, "swaggertestvault", "SwaggerTestRg", "newPolicyV2", armrecoveryservicesbackup.ProtectionPolicyResource{
		Properties: &armrecoveryservicesbackup.AzureFileShareProtectionPolicy{
			BackupManagementType: to.Ptr("AzureStorage"),
			SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
				SchedulePolicyType:   to.Ptr("SimpleSchedulePolicy"),
				ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
				ScheduleRunTimes: []*time.Time{
					to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t }())},
			},
			TimeZone: to.Ptr("UTC"),
			VaultRetentionPolicy: &armrecoveryservicesbackup.VaultRetentionPolicy{
				SnapshotRetentionInDays: to.Ptr[int32](5),
				VaultRetention: &armrecoveryservicesbackup.LongTermRetentionPolicy{
					RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
					DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](30),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t }())},
					},
					MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](60),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
						},
						RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
						RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
								to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
							WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
								to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst)},
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t }())},
					},
					WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
						DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](12),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t }())},
					},
					YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
						MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
							to.Ptr(armrecoveryservicesbackup.MonthOfYearJanuary)},
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](10),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
						},
						RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
						RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
								to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
							WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
								to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst)},
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t }())},
					},
				},
			},
			WorkLoadType: to.Ptr(armrecoveryservicesbackup.WorkloadTypeAzureFileShare),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionPolicyResource = armrecoveryservicesbackup.ProtectionPolicyResource{
	// 	Name: to.Ptr("newPolicyV2"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupPolicies/newPolicyV2"),
	// 	Properties: &armrecoveryservicesbackup.AzureFileShareProtectionPolicy{
	// 		BackupManagementType: to.Ptr("AzureStorage"),
	// 		ProtectedItemsCount: to.Ptr[int32](0),
	// 		SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
	// 			SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
	// 			ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
	// 			ScheduleRunTimes: []*time.Time{
	// 				to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t}())},
	// 			},
	// 			TimeZone: to.Ptr("UTC"),
	// 			VaultRetentionPolicy: &armrecoveryservicesbackup.VaultRetentionPolicy{
	// 				SnapshotRetentionInDays: to.Ptr[int32](5),
	// 				VaultRetention: &armrecoveryservicesbackup.LongTermRetentionPolicy{
	// 					RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
	// 					DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
	// 						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 							Count: to.Ptr[int32](30),
	// 							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
	// 						},
	// 						RetentionTimes: []*time.Time{
	// 							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t}())},
	// 						},
	// 						MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
	// 							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 								Count: to.Ptr[int32](60),
	// 								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
	// 							},
	// 							RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 							RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 								DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 									to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
	// 									WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 										to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst)},
	// 									},
	// 									RetentionTimes: []*time.Time{
	// 										to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t}())},
	// 									},
	// 									WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
	// 										DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 											to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
	// 											RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 												Count: to.Ptr[int32](12),
	// 												DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
	// 											},
	// 											RetentionTimes: []*time.Time{
	// 												to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t}())},
	// 											},
	// 											YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
	// 												MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
	// 													to.Ptr(armrecoveryservicesbackup.MonthOfYearJanuary)},
	// 													RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 														Count: to.Ptr[int32](10),
	// 														DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
	// 													},
	// 													RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 													RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 														DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 															to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
	// 															WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 																to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst)},
	// 															},
	// 															RetentionTimes: []*time.Time{
	// 																to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-18T09:30:00.000Z"); return t}())},
	// 															},
	// 														},
	// 													},
	// 												},
	// 											}
}
