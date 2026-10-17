
/**
 * Samples for EdgeMachineDiskPrivilegedJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskPrivilegedJobs_Get_DeleteVolume.json
     */
    /**
     * Sample code: EdgeMachineDiskPrivilegedJobs_Get_DeleteVolume.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDiskPrivilegedJobsGetDeleteVolume(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskPrivilegedJobs().getWithResponse("test-rg", "EdgeMachine01", "disk-001", "DeleteVolume",
            com.azure.core.util.Context.NONE);
    }
}
