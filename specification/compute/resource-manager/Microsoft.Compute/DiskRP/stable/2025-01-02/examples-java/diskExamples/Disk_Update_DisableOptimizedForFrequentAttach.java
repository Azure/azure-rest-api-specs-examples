
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/
     * Disk_Update_DisableOptimizedForFrequentAttach.json
     */
    /**
     * Sample code: update a managed disk to disable optimizedForFrequentAttach.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateAManagedDiskToDisableOptimizedForFrequentAttach(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withOptimizedForFrequentAttach(false), com.azure.core.util.Context.NONE);
    }
}
