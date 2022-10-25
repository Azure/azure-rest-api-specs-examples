import com.azure.core.util.Context;

/** Samples for BackupJobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/ListJobsWithStartTimeAndEndTimeFilters.json
     */
    /**
     * Sample code: List Jobs With Time Filter.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listJobsWithTimeFilter(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupJobs()
            .list(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "startTime eq '2016-01-01 00:00:00 AM' and endTime eq '2017-11-29 00:00:00 AM'",
                null,
                Context.NONE);
    }
}
