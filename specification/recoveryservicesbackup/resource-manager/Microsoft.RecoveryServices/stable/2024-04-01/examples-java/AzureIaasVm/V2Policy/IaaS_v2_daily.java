
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureIaaSvmProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.DailyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.DailySchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.DayOfWeek;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMSnapshotConsistencyType;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasvmPolicyType;
import com.azure.resourcemanager.recoveryservicesbackup.models.LongTermRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthOfYear;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthlyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDuration;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDurationType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionScheduleFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.ScheduleRunType;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleSchedulePolicyV2;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeekOfMonth;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.YearlyRetentionSchedule;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for ProtectionPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/V2Policy/IaaS_v2_daily.json
     */
    /**
     * Sample code: Create or Update Enhanced Azure Vm Protection Policy with daily backup.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateEnhancedAzureVmProtectionPolicyWithDailyBackup(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies().define(
                "v2-daily-sample")
            .withRegion((String) null).withExistingVault("NetSDKTestRsVault", "SwaggerTestRg")
            .withProperties(new AzureIaaSvmProtectionPolicy()
                .withSchedulePolicy(
                    new SimpleSchedulePolicyV2().withScheduleRunFrequency(ScheduleRunType.DAILY)
                        .withDailySchedule(new DailySchedule()
                            .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))))
                .withRetentionPolicy(new LongTermRetentionPolicy()
                    .withDailySchedule(new DailyRetentionSchedule()
                        .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-12-17T08:00:00+00:00")))
                        .withRetentionDuration(
                            new RetentionDuration().withCount(180).withDurationType(RetentionDurationType.DAYS)))
                    .withWeeklySchedule(new WeeklyRetentionSchedule().withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                        .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-12-17T08:00:00+00:00")))
                        .withRetentionDuration(
                            new RetentionDuration().withCount(12).withDurationType(RetentionDurationType.WEEKS)))
                    .withMonthlySchedule(
                        new MonthlyRetentionSchedule().withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                            .withRetentionScheduleWeekly(
                                new WeeklyRetentionFormat().withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FIRST)))
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-12-17T08:00:00+00:00")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(60).withDurationType(RetentionDurationType.MONTHS)))
                    .withYearlySchedule(
                        new YearlyRetentionSchedule().withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                            .withMonthsOfYear(Arrays.asList(MonthOfYear.JANUARY))
                            .withRetentionScheduleWeekly(
                                new WeeklyRetentionFormat().withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FIRST)))
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-12-17T08:00:00+00:00")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(10).withDurationType(RetentionDurationType.YEARS))))
                .withInstantRpRetentionRangeInDays(30).withTimeZone("India Standard Time")
                .withPolicyType(IaasvmPolicyType.V2)
                .withSnapshotConsistencyType(IaasVMSnapshotConsistencyType.ONLY_CRASH_CONSISTENT))
            .create();
    }
}
