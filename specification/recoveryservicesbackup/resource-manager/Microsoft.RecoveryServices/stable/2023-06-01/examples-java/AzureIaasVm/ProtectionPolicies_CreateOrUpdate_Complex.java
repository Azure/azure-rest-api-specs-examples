
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureIaaSvmProtectionPolicy;
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
import com.azure.resourcemanager.recoveryservicesbackup.models.YearlyRetentionSchedule;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for ProtectionPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Complex.json
     */
    /**
     * Sample code: Create or Update Full Azure Vm Protection Policy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateFullAzureVmProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().define("testPolicy1").withRegion((String) null)
            .withExistingVault("NetSDKTestRsVault", "SwaggerTestRg")
            .withProperties(
                new AzureIaaSvmProtectionPolicy()
                    .withSchedulePolicy(new SimpleSchedulePolicy().withScheduleRunFrequency(ScheduleRunType.WEEKLY)
                        .withScheduleRunDays(Arrays.asList(DayOfWeek.MONDAY, DayOfWeek.WEDNESDAY,
                            DayOfWeek.THURSDAY))
                        .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z"))))
                    .withRetentionPolicy(
                        new LongTermRetentionPolicy()
                            .withWeeklySchedule(
                                new WeeklyRetentionSchedule()
                                    .withDaysOfTheWeek(
                                        Arrays.asList(DayOfWeek.MONDAY, DayOfWeek.WEDNESDAY, DayOfWeek.THURSDAY))
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                    .withRetentionDuration(new RetentionDuration().withCount(1)
                                        .withDurationType(RetentionDurationType.WEEKS)))
                            .withMonthlySchedule(new MonthlyRetentionSchedule()
                                .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                .withRetentionScheduleWeekly(new WeeklyRetentionFormat()
                                    .withDaysOfTheWeek(Arrays.asList(DayOfWeek.WEDNESDAY, DayOfWeek.THURSDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FIRST, WeekOfMonth.THIRD)))
                                .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                .withRetentionDuration(new RetentionDuration().withCount(2)
                                    .withDurationType(RetentionDurationType.MONTHS)))
                            .withYearlySchedule(new YearlyRetentionSchedule()
                                .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                                .withMonthsOfYear(Arrays.asList(MonthOfYear.FEBRUARY, MonthOfYear.NOVEMBER))
                                .withRetentionScheduleWeekly(new WeeklyRetentionFormat()
                                    .withDaysOfTheWeek(Arrays.asList(DayOfWeek.MONDAY, DayOfWeek.THURSDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.FOURTH)))
                                .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T10:00:00Z")))
                                .withRetentionDuration(new RetentionDuration().withCount(4)
                                    .withDurationType(RetentionDurationType.YEARS))))
                    .withTimeZone("Pacific Standard Time"))
            .create();
    }
}
