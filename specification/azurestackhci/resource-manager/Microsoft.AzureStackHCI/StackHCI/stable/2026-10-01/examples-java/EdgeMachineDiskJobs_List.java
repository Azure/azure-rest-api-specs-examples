
/**
 * Samples for EdgeMachineDiskJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskJobs_List.json
     */
    /**
     * Sample code: EdgeMachineDiskJobs_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDiskJobsList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskJobs().list("test-rg", "EdgeMachine01", "disk-001", com.azure.core.util.Context.NONE);
    }
}
