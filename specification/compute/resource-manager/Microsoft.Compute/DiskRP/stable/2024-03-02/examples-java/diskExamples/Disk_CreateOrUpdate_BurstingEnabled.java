
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_CreateOrUpdate_BurstingEnabled.json
     */
    /**
     * Sample code: Create or update a bursting enabled managed disk.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateABurstingEnabledManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withDiskSizeGB(1024).withBurstingEnabled(true), com.azure.core.util.Context.NONE);
    }
}
