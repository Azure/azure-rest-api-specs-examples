
/**
 * Samples for DiskRestorePoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get_WithInstantAccess.json
     */
    /**
     * Sample code: get a disk restorePoint resource with InstantAccess snapshotAccessState.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getADiskRestorePointResourceWithInstantAccessSnapshotAccessState(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().getWithResponse("myResourceGroup", "rpc", "vmrp",
            "TestDisk45ceb03433006d1baee_5c1528-43e2-4c77-9c55-a78bf5a5fc88", com.azure.core.util.Context.NONE);
    }
}
