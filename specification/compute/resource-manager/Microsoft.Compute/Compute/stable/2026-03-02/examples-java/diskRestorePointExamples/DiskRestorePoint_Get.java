
/**
 * Samples for DiskRestorePoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get.json
     */
    /**
     * Sample code: get an incremental disk restorePoint resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAnIncrementalDiskRestorePointResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().getWithResponse("myResourceGroup", "rpc", "vmrp",
            "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", com.azure.core.util.Context.NONE);
    }
}
