
/**
 * Samples for DiskRestorePoint RevokeAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskRestorePointExamples/DiskRestorePoint_EndGetAccess.json
     */
    /**
     * Sample code: revokes access to a diskRestorePoint.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void revokesAccessToADiskRestorePoint(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().revokeAccess("myResourceGroup", "rpc", "vmrp",
            "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", com.azure.core.util.Context.NONE);
    }
}
