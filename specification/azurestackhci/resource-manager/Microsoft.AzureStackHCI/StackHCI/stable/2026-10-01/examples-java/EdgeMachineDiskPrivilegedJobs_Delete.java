
/**
 * Samples for EdgeMachineDiskPrivilegedJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskPrivilegedJobs_Delete.json
     */
    /**
     * Sample code: EdgeMachineDiskPrivilegedJobs_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineDiskPrivilegedJobsDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskPrivilegedJobs().delete("test-rg", "EdgeMachine01", "disk-001", "DeleteVolume",
            com.azure.core.util.Context.NONE);
    }
}
