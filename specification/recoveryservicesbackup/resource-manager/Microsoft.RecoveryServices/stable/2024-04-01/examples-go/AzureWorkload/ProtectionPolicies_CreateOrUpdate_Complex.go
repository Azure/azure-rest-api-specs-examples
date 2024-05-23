package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureWorkload/ProtectionPolicies_CreateOrUpdate_Complex.json
func ExampleProtectionPoliciesClient_CreateOrUpdate_createOrUpdateFullAzureWorkloadProtectionPolicy() {
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
		Properties: &armrecoveryservicesbackup.AzureVMWorkloadProtectionPolicy{
			BackupManagementType: to.Ptr("AzureWorkload"),
			Settings: &armrecoveryservicesbackup.Settings{
				Issqlcompression: to.Ptr(false),
				TimeZone:         to.Ptr("Pacific Standard Time"),
			},
			SubProtectionPolicy: []*armrecoveryservicesbackup.SubProtectionPolicy{
				{
					PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeFull),
					RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
						RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
						MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
								Count:        to.Ptr[int32](1),
								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
							},
							RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
							RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
								DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
									to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
								WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
									to.Ptr(armrecoveryservicesbackup.WeekOfMonthSecond)},
							},
							RetentionTimes: []*time.Time{
								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
						},
						WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
								to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
								to.Ptr(armrecoveryservicesbackup.DayOfWeekTuesday)},
							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
								Count:        to.Ptr[int32](2),
								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
							},
							RetentionTimes: []*time.Time{
								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
						},
						YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
							MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
								to.Ptr(armrecoveryservicesbackup.MonthOfYearJanuary),
								to.Ptr(armrecoveryservicesbackup.MonthOfYearJune),
								to.Ptr(armrecoveryservicesbackup.MonthOfYearDecember)},
							RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
								Count:        to.Ptr[int32](1),
								DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
							},
							RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
							RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
								DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
									to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
								WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
									to.Ptr(armrecoveryservicesbackup.WeekOfMonthLast)},
							},
							RetentionTimes: []*time.Time{
								to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
						},
					},
					SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
						SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
						ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
							to.Ptr(armrecoveryservicesbackup.DayOfWeekTuesday)},
						ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
						ScheduleRunTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
					},
				},
				{
					PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeDifferential),
					RetentionPolicy: &armrecoveryservicesbackup.SimpleRetentionPolicy{
						RetentionPolicyType: to.Ptr("SimpleRetentionPolicy"),
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](8),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
						},
					},
					SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
						SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
						ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekFriday)},
						ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
						ScheduleRunTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t }())},
					},
				},
				{
					PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeLog),
					RetentionPolicy: &armrecoveryservicesbackup.SimpleRetentionPolicy{
						RetentionPolicyType: to.Ptr("SimpleRetentionPolicy"),
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](7),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
						},
					},
					SchedulePolicy: &armrecoveryservicesbackup.LogSchedulePolicy{
						SchedulePolicyType:      to.Ptr("LogSchedulePolicy"),
						ScheduleFrequencyInMins: to.Ptr[int32](60),
					},
				}},
			WorkLoadType: to.Ptr(armrecoveryservicesbackup.WorkloadTypeSQLDataBase),
		},
	}, &armrecoveryservicesbackup.ProtectionPoliciesClientCreateOrUpdateOptions{XMSAuthorizationAuxiliary: nil})
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
	// 	Properties: &armrecoveryservicesbackup.AzureVMWorkloadProtectionPolicy{
	// 		BackupManagementType: to.Ptr("AzureWorkload"),
	// 		ProtectedItemsCount: to.Ptr[int32](0),
	// 		Settings: &armrecoveryservicesbackup.Settings{
	// 			Issqlcompression: to.Ptr(false),
	// 			TimeZone: to.Ptr("Pacific Standard Time"),
	// 		},
	// 		SubProtectionPolicy: []*armrecoveryservicesbackup.SubProtectionPolicy{
	// 			{
	// 				PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeFull),
	// 				RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
	// 					RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
	// 					MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
	// 						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 							Count: to.Ptr[int32](1),
	// 							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeMonths),
	// 						},
	// 						RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 						RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 								to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
	// 								WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 									to.Ptr(armrecoveryservicesbackup.WeekOfMonthSecond)},
	// 								},
	// 								RetentionTimes: []*time.Time{
	// 									to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 								},
	// 								WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
	// 									DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 										to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
	// 										to.Ptr(armrecoveryservicesbackup.DayOfWeekTuesday)},
	// 										RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 											Count: to.Ptr[int32](2),
	// 											DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
	// 										},
	// 										RetentionTimes: []*time.Time{
	// 											to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 										},
	// 										YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
	// 											MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
	// 												to.Ptr(armrecoveryservicesbackup.MonthOfYearJanuary),
	// 												to.Ptr(armrecoveryservicesbackup.MonthOfYearJune),
	// 												to.Ptr(armrecoveryservicesbackup.MonthOfYearDecember)},
	// 												RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 													Count: to.Ptr[int32](1),
	// 													DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeYears),
	// 												},
	// 												RetentionScheduleFormatType: to.Ptr(armrecoveryservicesbackup.RetentionScheduleFormatWeekly),
	// 												RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
	// 													DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
	// 														to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
	// 														WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
	// 															to.Ptr(armrecoveryservicesbackup.WeekOfMonthLast)},
	// 														},
	// 														RetentionTimes: []*time.Time{
	// 															to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 														},
	// 													},
	// 													SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
	// 														SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
	// 														ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
	// 															to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday),
	// 															to.Ptr(armrecoveryservicesbackup.DayOfWeekTuesday)},
	// 															ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
	// 															ScheduleRunTimes: []*time.Time{
	// 																to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 																ScheduleWeeklyFrequency: to.Ptr[int32](0),
	// 															},
	// 														},
	// 														{
	// 															PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeDifferential),
	// 															RetentionPolicy: &armrecoveryservicesbackup.SimpleRetentionPolicy{
	// 																RetentionPolicyType: to.Ptr("SimpleRetentionPolicy"),
	// 																RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 																	Count: to.Ptr[int32](8),
	// 																	DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
	// 																},
	// 															},
	// 															SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
	// 																SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
	// 																ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
	// 																	to.Ptr(armrecoveryservicesbackup.DayOfWeekFriday)},
	// 																	ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeWeekly),
	// 																	ScheduleRunTimes: []*time.Time{
	// 																		to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00.000Z"); return t}())},
	// 																		ScheduleWeeklyFrequency: to.Ptr[int32](0),
	// 																	},
	// 																},
	// 																{
	// 																	PolicyType: to.Ptr(armrecoveryservicesbackup.PolicyTypeLog),
	// 																	RetentionPolicy: &armrecoveryservicesbackup.SimpleRetentionPolicy{
	// 																		RetentionPolicyType: to.Ptr("SimpleRetentionPolicy"),
	// 																		RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
	// 																			Count: to.Ptr[int32](7),
	// 																			DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
	// 																		},
	// 																	},
	// 																	SchedulePolicy: &armrecoveryservicesbackup.LogSchedulePolicy{
	// 																		SchedulePolicyType: to.Ptr("LogSchedulePolicy"),
	// 																		ScheduleFrequencyInMins: to.Ptr[int32](60),
	// 																	},
	// 															}},
	// 															WorkLoadType: to.Ptr(armrecoveryservicesbackup.WorkloadTypeSQLDataBase),
	// 														},
	// 													}
}
