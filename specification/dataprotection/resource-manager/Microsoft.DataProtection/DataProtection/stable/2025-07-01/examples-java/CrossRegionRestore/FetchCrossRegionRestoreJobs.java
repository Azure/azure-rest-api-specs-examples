
import com.azure.resourcemanager.dataprotection.models.CrossRegionRestoreJobsRequest;

/**
 * Samples for FetchCrossRegionRestoreJobsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CrossRegionRestore/FetchCrossRegionRestoreJobs.json
     */
    /**
     * Sample code: List Cross Region Restore Jobs.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listCrossRegionRestoreJobs(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.fetchCrossRegionRestoreJobsOperations().list("BugBash1", "east us",
            new CrossRegionRestoreJobsRequest().withSourceRegion("east us").withSourceBackupVaultId(
                "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/backupVaults/BugBashVaultForCCYv11"),
            null, com.azure.core.util.Context.NONE);
    }
}
