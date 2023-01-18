/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/JobCRUD/ListJobs.json
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
