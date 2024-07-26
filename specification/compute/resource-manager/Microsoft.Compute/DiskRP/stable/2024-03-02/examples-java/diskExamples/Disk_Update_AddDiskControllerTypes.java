
import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.SupportedCapabilities;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Update_AddDiskControllerTypes.json
     */
    /**
     * Sample code: Update a managed disk with diskControllerTypes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAManagedDiskWithDiskControllerTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withSupportedCapabilities(new SupportedCapabilities().withDiskControllerTypes("SCSI")),
            com.azure.core.util.Context.NONE);
    }
}
