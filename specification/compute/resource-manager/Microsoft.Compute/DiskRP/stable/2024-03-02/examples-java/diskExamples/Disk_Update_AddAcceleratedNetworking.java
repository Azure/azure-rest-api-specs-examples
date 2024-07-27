
import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.SupportedCapabilities;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Update_AddAcceleratedNetworking.json
     */
    /**
     * Sample code: Update a managed disk to add accelerated networking.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateAManagedDiskToAddAcceleratedNetworking(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withSupportedCapabilities(new SupportedCapabilities().withAcceleratedNetwork(false)),
            com.azure.core.util.Context.NONE);
    }
}
