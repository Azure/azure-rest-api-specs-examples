import com.azure.resourcemanager.recoveryservicesbackup.models.AzureFileShareProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.DailyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.DayOfWeek;
import com.azure.resourcemanager.recoveryservicesbackup.models.LongTermRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthOfYear;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthlyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDuration;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDurationType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionScheduleFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.ScheduleRunType;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleSchedulePolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeekOfMonth;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.WeeklyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.WorkloadType;
import com.azure.resourcemanager.recoveryservicesbackup.models.YearlyRetentionSchedule;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for ProtectionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Daily.json
     */
    /**
     * Sample code: Create or Update Daily Azure Storage Protection Policy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateDailyAzureStorageProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies()
            .define("dailyPolicy2")
            .withRegion((String) null)
            .withExistingVault("swaggertestvault", "SwaggerTestRg")
            .withProperties(
                new AzureFileShareProtectionPolicy()
                    .withWorkLoadType(WorkloadType.AZURE_FILE_SHARE)
                    .withSchedulePolicy(
                        new SimpleSchedulePolicy()
                            .withScheduleRunFrequency(ScheduleRunType.DAILY)
                            .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2021-09-29T08:00:00.000Z"))))
                    .withRetentionPolicy(
                        new LongTermRetentionPolicy()
                            .withDailySchedule(
                                new DailyRetentionSchedule()
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-09-29T08:00:00.000Z")))
                                    .withRetentionDuration(
                                        new RetentionDuration()
                                            .withCount(5)
                                            .withDurationType(RetentionDurationType.DAYS)))
                            .withWeeklySchedule(
                                new WeeklyRetentionSchedule()
                                    .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-09-29T08:00:00.000Z")))
                                    .withRetentionDuration(
                                        new RetentionDuration()
                                            .withCount(12)
                                            .withDurationType(RetentionDurationType.WEEKS)))
                            .withMonthlySchedule(
                                new MonthlyRetentionSchedule()
                                    .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                    .withRetentionScheduleWeekly(
                                        new WeeklyRetentionFormat()
                                            .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                            .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FIRST)))
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-09-29T08:00:00.000Z")))
                                    .withRetentionDuration(
                                        new RetentionDuration()
                                            .withCount(60)
                                            .withDurationType(RetentionDurationType.MONTHS)))
                            .withYearlySchedule(
                                new YearlyRetentionSchedule()
                                    .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                    .withMonthsOfYear(Arrays.asList(MonthOfYear.JANUARY))
                                    .withRetentionScheduleWeekly(
                                        new WeeklyRetentionFormat()
                                            .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                            .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FIRST)))
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2021-09-29T08:00:00.000Z")))
                                    .withRetentionDuration(
                                        new RetentionDuration()
                                            .withCount(10)
                                            .withDurationType(RetentionDurationType.YEARS))))
                    .withTimeZone("UTC"))
            .create();
    }
}
