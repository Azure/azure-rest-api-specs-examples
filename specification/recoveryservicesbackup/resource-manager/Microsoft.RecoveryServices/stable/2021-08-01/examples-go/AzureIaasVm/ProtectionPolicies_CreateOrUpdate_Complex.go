package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Complex.json
func ExampleProtectionPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<policy-name>",
		armrecoveryservicesbackup.ProtectionPolicyResource{
			Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
				ProtectionPolicy: armrecoveryservicesbackup.ProtectionPolicy{
					BackupManagementType: to.StringPtr("<backup-management-type>"),
				},
				RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
					RetentionPolicy: armrecoveryservicesbackup.RetentionPolicy{
						RetentionPolicyType: to.StringPtr("<retention-policy-type>"),
					},
					MonthlySchedule: &armrecoveryservicesbackup.MonthlyRetentionSchedule{
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Int32Ptr(2),
							DurationType: armrecoveryservicesbackup.RetentionDurationTypeMonths.ToPtr(),
						},
						RetentionScheduleFormatType: armrecoveryservicesbackup.RetentionScheduleFormatWeekly.ToPtr(),
						RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
								armrecoveryservicesbackup.DayOfWeekWednesday.ToPtr(),
								armrecoveryservicesbackup.DayOfWeekThursday.ToPtr()},
							WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
								armrecoveryservicesbackup.WeekOfMonthFirst.ToPtr(),
								armrecoveryservicesbackup.WeekOfMonthThird.ToPtr()},
						},
						RetentionTimes: []*time.Time{
							to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00Z"); return t }())},
					},
					WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
						DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
							armrecoveryservicesbackup.DayOfWeekMonday.ToPtr(),
							armrecoveryservicesbackup.DayOfWeekWednesday.ToPtr(),
							armrecoveryservicesbackup.DayOfWeekThursday.ToPtr()},
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Int32Ptr(1),
							DurationType: armrecoveryservicesbackup.RetentionDurationTypeWeeks.ToPtr(),
						},
						RetentionTimes: []*time.Time{
							to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00Z"); return t }())},
					},
					YearlySchedule: &armrecoveryservicesbackup.YearlyRetentionSchedule{
						MonthsOfYear: []*armrecoveryservicesbackup.MonthOfYear{
							armrecoveryservicesbackup.MonthOfYearFebruary.ToPtr(),
							armrecoveryservicesbackup.MonthOfYearNovember.ToPtr()},
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Int32Ptr(4),
							DurationType: armrecoveryservicesbackup.RetentionDurationTypeYears.ToPtr(),
						},
						RetentionScheduleFormatType: armrecoveryservicesbackup.RetentionScheduleFormatWeekly.ToPtr(),
						RetentionScheduleWeekly: &armrecoveryservicesbackup.WeeklyRetentionFormat{
							DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
								armrecoveryservicesbackup.DayOfWeekMonday.ToPtr(),
								armrecoveryservicesbackup.DayOfWeekThursday.ToPtr()},
							WeeksOfTheMonth: []*armrecoveryservicesbackup.WeekOfMonth{
								armrecoveryservicesbackup.WeekOfMonthFourth.ToPtr()},
						},
						RetentionTimes: []*time.Time{
							to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00Z"); return t }())},
					},
				},
				SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
					SchedulePolicy: armrecoveryservicesbackup.SchedulePolicy{
						SchedulePolicyType: to.StringPtr("<schedule-policy-type>"),
					},
					ScheduleRunDays: []*armrecoveryservicesbackup.DayOfWeek{
						armrecoveryservicesbackup.DayOfWeekMonday.ToPtr(),
						armrecoveryservicesbackup.DayOfWeekWednesday.ToPtr(),
						armrecoveryservicesbackup.DayOfWeekThursday.ToPtr()},
					ScheduleRunFrequency: armrecoveryservicesbackup.ScheduleRunTypeWeekly.ToPtr(),
					ScheduleRunTimes: []*time.Time{
						to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T10:00:00Z"); return t }())},
				},
				TimeZone: to.StringPtr("<time-zone>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProtectionPolicyResource.ID: %s\n", *res.ID)
}
