
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/JobCRUD/GetJob.
     * json
     */
    /**
     * Sample code: Get Job.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getJob(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.jobs().getWithResponse("BugBash1", "BugBashVaultForCCYv11", "3c60cb49-63e8-4b21-b9bd-26277b3fdfae",
            com.azure.core.util.Context.NONE);
    }
}
