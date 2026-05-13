
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/JobCRUD/ListJobs.json
     */
    /**
     * Sample code: Get Jobs.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getJobs(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.jobs().list("BugBash1", "BugBashVaultForCCYv11", com.azure.core.util.Context.NONE);
    }
}
