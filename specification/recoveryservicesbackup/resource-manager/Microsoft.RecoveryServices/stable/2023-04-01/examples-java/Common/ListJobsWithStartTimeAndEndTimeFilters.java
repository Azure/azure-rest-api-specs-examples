/** Samples for BackupJobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/ListJobsWithStartTimeAndEndTimeFilters.json
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
                com.azure.core.util.Context.NONE);
    }
}
