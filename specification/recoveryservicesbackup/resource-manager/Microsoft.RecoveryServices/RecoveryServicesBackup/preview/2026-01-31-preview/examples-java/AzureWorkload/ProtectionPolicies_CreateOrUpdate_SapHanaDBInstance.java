
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureVmWorkloadProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.DailyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.DayOfWeek;
import com.azure.resourcemanager.recoveryservicesbackup.models.LongTermRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthOfYear;
import com.azure.resourcemanager.recoveryservicesbackup.models.MonthlyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.PolicyType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDuration;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDurationType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionScheduleFormat;
import com.azure.resourcemanager.recoveryservicesbackup.models.ScheduleRunType;
import com.azure.resourcemanager.recoveryservicesbackup.models.Settings;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleSchedulePolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.SnapshotBackupAdditionalDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.SubProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.UserAssignedIdentityProperties;
import com.azure.resourcemanager.recoveryservicesbackup.models.UserAssignedManagedIdentityDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.VMWorkloadPolicyType;
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
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/ProtectionPolicies_CreateOrUpdate_SapHanaDBInstance.json
     */
    /**
     * Sample code: Create or Update Sap Hana DB Instance Workload Protection Policy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateSapHanaDBInstanceWorkloadProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies().define(
                "testHanaSnapshotV2Policy1")
            .withExistingVault("HanaTestRsVault",
                "SwaggerTestRg")
            .withProperties(new AzureVmWorkloadProtectionPolicy().withProtectedItemsCount(0)
                .withWorkLoadType(WorkloadType.SAPHANA_DBINSTANCE)
                .withVmWorkloadPolicyType(VMWorkloadPolicyType.SNAPSHOT_V2)
                .withSettings(new Settings().withTimeZone("UTC").withIssqlcompression(false).withIsCompression(false))
                .withSubProtectionPolicy(Arrays.asList(new SubProtectionPolicy()
                    .withPolicyType(PolicyType.SNAPSHOT_FULL)
                    .withSchedulePolicy(new SimpleSchedulePolicy().withScheduleRunFrequency(ScheduleRunType.DAILY)
                        .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2024-10-01T03:30:00.000Z"))))
                    .withRetentionPolicy(new LongTermRetentionPolicy()
                        .withDailySchedule(new DailyRetentionSchedule()
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2023-12-19T20:00:00.000Z")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(30).withDurationType(RetentionDurationType.DAYS)))
                        .withWeeklySchedule(new WeeklyRetentionSchedule()
                            .withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2023-12-19T20:00:00.000Z")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(10).withDurationType(RetentionDurationType.WEEKS)))
                        .withMonthlySchedule(new MonthlyRetentionSchedule()
                            .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                            .withRetentionScheduleWeekly(
                                new WeeklyRetentionFormat().withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.SECOND)))
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2023-12-15T20:00:00.000Z")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(6).withDurationType(RetentionDurationType.MONTHS)))
                        .withYearlySchedule(new YearlyRetentionSchedule()
                            .withRetentionScheduleFormatType(RetentionScheduleFormat.WEEKLY)
                            .withMonthsOfYear(Arrays.asList(MonthOfYear.JANUARY))
                            .withRetentionScheduleWeekly(
                                new WeeklyRetentionFormat().withDaysOfTheWeek(Arrays.asList(DayOfWeek.SUNDAY))
                                    .withWeeksOfTheMonth(Arrays.asList(WeekOfMonth.LAST)))
                            .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2023-12-19T20:00:00.000Z")))
                            .withRetentionDuration(
                                new RetentionDuration().withCount(2).withDurationType(RetentionDurationType.YEARS))))
                    .withSnapshotBackupAdditionalDetails(new SnapshotBackupAdditionalDetails()
                        .withInstantRpRetentionRangeInDays(5).withInstantRPDetails("SwaggerTestRG")
                        .withUserAssignedManagedIdentityDetails(new UserAssignedManagedIdentityDetails()
                            .withIdentityArmId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerMsiRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/SwaggerUMI")
                            .withIdentityName("SwaggerUMI")
                            .withUserAssignedIdentityProperties(new UserAssignedIdentityProperties()
                                .withClientId("00000000-0000-0000-0000-000000000000")
                                .withPrincipalId("00000000-0000-0000-0000-000000000000")))))))
            .create();
    }
}
