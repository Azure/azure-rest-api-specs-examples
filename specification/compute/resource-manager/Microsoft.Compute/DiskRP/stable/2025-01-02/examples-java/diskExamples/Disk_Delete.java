
/**
 * Samples for Disks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/
     * Disk_Delete.json
     */
    /**
     * Sample code: delete a managed disk.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().delete("myResourceGroup", "myDisk",
            com.azure.core.util.Context.NONE);
    }
}
