
/**
 * Samples for DiskRestorePoint ListByRestorePoint.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/
     * diskRestorePointExamples/DiskRestorePoint_ListByVmRestorePoint.json
     */
    /**
     * Sample code: Get an incremental disk restorePoint resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnIncrementalDiskRestorePointResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskRestorePoints().listByRestorePoint("myResourceGroup",
            "rpc", "vmrp", com.azure.core.util.Context.NONE);
    }
}
