
/**
 * Samples for DiskRestorePoint ListByRestorePoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskRestorePointExamples/DiskRestorePoint_ListByVmRestorePoint.json
     */
    /**
     * Sample code: get an incremental disk restorePoint resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAnIncrementalDiskRestorePointResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().listByRestorePoint("myResourceGroup", "rpc", "vmrp",
            com.azure.core.util.Context.NONE);
    }
}
