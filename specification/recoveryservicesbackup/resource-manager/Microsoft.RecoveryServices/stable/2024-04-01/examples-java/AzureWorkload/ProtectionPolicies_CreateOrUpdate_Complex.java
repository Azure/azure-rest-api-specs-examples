
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureVmWorkloadProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.DayOfWeek;
import com.azure.resourcemanager.recoveryservicesbackup.models.LogSchedulePolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.LongTermRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthOfYear;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthlyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.PolicyType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDuration;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDurationType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionScheduleFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.ScheduleRunType;
import com.azure.resourcemanager.recoveryservicesbackup.models.Settings;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleSchedulePolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.SubProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeekOfMonth;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.WorkloadType;
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
     * AzureWorkload/ProtectionPolicies_CreateOrUpdate_Complex.json
     */
    /**
     * Sample code: Create or Update Full Azure Workload Protection Policy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateFullAzureWorkloadProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies().define(
                "testPolicy1")
            .withRegion((String) null).withExistingVault("NetSDKTestRsVault", "SwaggerTestRg")
            .withProperties(
                new AzureVmWorkloadProtectionPolicy().withWorkLoadType(WorkloadType.SQLDATA_BASE)
                    .withSettings(new Settings().withTimeZone("Pacific Standard Time").withIssqlcompression(false))
                    .withSubProtectionPolicy(Arrays.asList(
                        new SubProtectionPolicy().withPolicyType(PolicyType.FULL)
                            .withSchedulePolicy(
                                new SimpleSchedulePolicy().withScheduleRunFrequency(ScheduleRunType.WEEKLY)
                                    .withScheduleRunDays(Arrays.asList(DayOfWeek.SUNDAY, DayOfWeek.TUESDAY))
                                    .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z"))))
                            .withRetentionPolicy(
                                new LongTermRetentionPolicy()
                                    .withWeeklySchedule(new WeeklyRetentionSchedule()
                                        .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY, DayOfWeek.TUESDAY))
                                        .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                        .withRetentionDuration(new RetentionDuration().withCount(2)
                                            .withDurationType(RetentionDurationType.WEEKS)))
                                    .withMonthlySchedule(new MonthlyRetentionSchedule()
                                        .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                        .withRetentionScheduleWeekly(new WeeklyRetentionFormat()
                                            .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                            .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.SECOND)))
                                        .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                        .withRetentionDuration(new RetentionDuration().withCount(1)
                                            .withDurationType(RetentionDurationType.MONTHS)))
                                    .withYearlySchedule(new YearlyRetentionSchedule()
                                        .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                        .withMonthsOfYear(
                                            Arrays.asList(MonthOfYear.JANUARY, MonthOfYear.JUNE, MonthOfYear.DECEMBER))
                                        .withRetentionScheduleWeekly(new WeeklyRetentionFormat()
                                            .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                            .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.LAST)))
                                        .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                        .withRetentionDuration(new RetentionDuration().withCount(1)
                                            .withDurationType(RetentionDurationType.YEARS)))),
                        new SubProtectionPolicy().withPolicyType(PolicyType.DIFFERENTIAL)
                            .withSchedulePolicy(
                                new SimpleSchedulePolicy().withScheduleRunFrequency(ScheduleRunType.WEEKLY)
                                    .withScheduleRunDays(Arrays.asList(DayOfWeek.FRIDAY))
                                    .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z"))))
                            .withRetentionPolicy(new SimpleRetentionPolicy().withRetentionDuration(
                                new RetentionDuration().withCount(8).withDurationType(RetentionDurationType.DAYS))),
                        new SubProtectionPolicy().withPolicyType(PolicyType.LOG)
                            .withSchedulePolicy(new LogSchedulePolicy().withScheduleFrequencyInMins(60))
                            .withRetentionPolicy(new SimpleRetentionPolicy().withRetentionDuration(
                                new RetentionDuration().withCount(7).withDurationType(RetentionDurationType.DAYS))))))
            .create();
    }
}
