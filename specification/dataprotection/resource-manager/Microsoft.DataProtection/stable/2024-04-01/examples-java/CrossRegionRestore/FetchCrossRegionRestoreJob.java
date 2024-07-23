
import com.azure.resourcemanager.dataprotection.models.CrossRegionRestoreJobRequest;

/**
 * Samples for FetchCrossRegionRestoreJob Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * CrossRegionRestore/FetchCrossRegionRestoreJob.json
     */
    /**
     * Sample code: Get Cross Region Restore Job.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getCrossRegionRestoreJob(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.fetchCrossRegionRestoreJobs().getWithResponse("BugBash1", "west us",
            new CrossRegionRestoreJobRequest().withSourceRegion("east us").withSourceBackupVaultId(
                "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/backupVaults/BugBashVaultForCCYv11")
                .withJobId("3c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
            com.azure.core.util.Context.NONE);
    }
}
