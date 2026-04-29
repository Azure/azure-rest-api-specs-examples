
/**
 * Samples for BackupJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/ListJobsWithAllSupportedFilters.json
     */
    /**
     * Sample code: List Jobs With Filters.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        listJobsWithFilters(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupJobs().list("NetSDKTestRsVault", "SwaggerTestRg",
            "startTime eq '2016-01-01 00:00:00 AM' and endTime eq '2017-11-29 00:00:00 AM' and operation eq 'Backup' and backupManagementType eq 'AzureIaasVM' and status eq 'InProgress'",
            null, com.azure.core.util.Context.NONE);
    }
}
