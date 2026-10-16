
/**
 * Samples for EdgeMachineDiskPrivilegedJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskPrivilegedJobs_List.json
     */
    /**
     * Sample code: EdgeMachineDiskPrivilegedJobs_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineDiskPrivilegedJobsList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskPrivilegedJobs().list("test-rg", "EdgeMachine01", "disk-001",
            com.azure.core.util.Context.NONE);
    }
}
