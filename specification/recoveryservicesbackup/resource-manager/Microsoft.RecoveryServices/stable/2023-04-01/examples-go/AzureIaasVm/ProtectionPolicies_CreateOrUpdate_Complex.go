package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a4ddec441435d1ef766c4f160eda658a69cc5dc2/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Complex.json
func ExampleProtectionPoliciesClient_CreateOrUpdate_createOrUpdateFullAzureVmProtectionPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionPoliciesClient().CreateOrUpdate(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1", armrecoveryservicesbackup.ProtectionPolicyResource{
		Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
			BackupManagementType: to.Ptr("AzureIaasVM"),
			RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
				RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
				MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
					RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
						Count:        to.Ptr[int32](2),
						DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
					},
					RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
					RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
						DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
							to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
						WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
							to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst),
							to.Ptr(armrecoveryservicesbackup.WeekOfMonthThird)},
					},
					RetentionTimes: []*time.Time{
						to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
				},
				WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
					DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
						to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
						to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
						to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
					RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
						Count:        to.Ptr[int32](1),
						DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
					},
					RetentionTimes: []*time.Time{
						to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
				},
				YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
					MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
						to.Ptr(armrecoveryservicesbackup.MonthOfYearFebruary),
						to.Ptr(armrecoveryservicesbackup.MonthOfYearNovember)},
					RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
						Count:        to.Ptr[int32](4),
						DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
					},
					RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
					RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
						DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
							to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
						WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
							to.Ptr(armrecoveryservicesbackup.WeekOfMonthFourth)},
					},
					RetentionTimes: []*time.Time{
						to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
				},
			},
			SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
				SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
				ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
					to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
					to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
					to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
				ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
				ScheduleRunTimes: []*time.Time{
					to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
			},
			TimeZone: to.Ptr("Pacific Standard Time"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionPolicyResource = armrecoveryservicesbackup.ProtectionPolicyResource{
	// 	Name: to.Ptr("testPolicy1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/testPolicy1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
	// 		BackupManagementType: to.Ptr("AzureIaasVM"),
	// 		ProtectedItemsCount: to.Ptr[int32](0),
	// 		RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
	// 			RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
	// 			MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
	// 				RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 					Count: to.Ptr[int32](2),
	// 					DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
	// 				},
	// 				RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 				RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 					DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 						to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
	// 						to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
	// 						WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 							to.Ptr(armrecoveryservicesbackup.WeekOfMonthFirst),
	// 							to.Ptr(armrecoveryservicesbackup.WeekOfMonthThird)},
	// 						},
	// 						RetentionTimes: []*time.Time{
	// 							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 						},
	// 						WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
	// 							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 								to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
	// 								to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
	// 								to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
	// 								RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 									Count: to.Ptr[int32](1),
	// 									DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
	// 								},
	// 								RetentionTimes: []*time.Time{
	// 									to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 								},
	// 								YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
	// 									MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
	// 										to.Ptr(armrecoveryservicesbackup.MonthOfYearFebruary),
	// 										to.Ptr(armrecoveryservicesbackup.MonthOfYearNovember)},
	// 										RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 											Count: to.Ptr[int32](4),
	// 											DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
	// 										},
	// 										RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 										RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 											DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 												to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
	// 												to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
	// 												WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 													to.Ptr(armrecoveryservicesbackup.WeekOfMonthFourth)},
	// 												},
	// 												RetentionTimes: []*time.Time{
	// 													to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 												},
	// 											},
	// 											SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
	// 												SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
	// 												ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
	// 													to.Ptr(armrecoveryservicesbackup.DayOfWeekMonday),
	// 													to.Ptr(armrecoveryservicesbackup.DayOfWeekWednesday),
	// 													to.Ptr(armrecoveryservicesbackup.DayOfWeekThursday)},
	// 													ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
	// 													ScheduleRunTimes: []*time.Time{
	// 														to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 														ScheduleWeeklyFrequency: to.Ptr[int32](0),
	// 													},
	// 													TimeZone: to.Ptr("Pacific Standard Time"),
	// 												},
	// 											}
}
