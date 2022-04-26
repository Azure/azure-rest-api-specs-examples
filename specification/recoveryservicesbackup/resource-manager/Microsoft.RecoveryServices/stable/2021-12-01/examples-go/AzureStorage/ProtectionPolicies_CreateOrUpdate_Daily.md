Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.5.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Daily.json
func ExampleProtectionPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectionPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<policy-name>",
		armrecoveryservicesbackup.ProtectionPolicyResource{
			Properties: &armrecoveryservicesbackup.AzureFileShareProtectionPolicy{
				BackupManagementType: to.Ptr("<backup-management-type>"),
				RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
					RetentionPolicyType: to.Ptr("<retention-policy-type>"),
					DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](5),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-29T08:00:00.000Z"); return t }())},
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
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-29T08:00:00.000Z"); return t }())},
					},
					WeeklySchedule: &armrecoveryservicesbackup.WeeklyRetentionSchedule{
						DaysOfTheWeek: []*armrecoveryservicesbackup.DayOfWeek{
							to.Ptr(armrecoveryservicesbackup.DayOfWeekSunday)},
						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
							Count:        to.Ptr[int32](12),
							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeWeeks),
						},
						RetentionTimes: []*time.Time{
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-29T08:00:00.000Z"); return t }())},
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
							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-29T08:00:00.000Z"); return t }())},
					},
				},
				SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
					SchedulePolicyType:   to.Ptr("<schedule-policy-type>"),
					ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
					ScheduleRunTimes: []*time.Time{
						to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-29T08:00:00.000Z"); return t }())},
				},
				TimeZone:     to.Ptr("<time-zone>"),
				WorkLoadType: to.Ptr(armrecoveryservicesbackup.WorkloadTypeAzureFileShare),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
