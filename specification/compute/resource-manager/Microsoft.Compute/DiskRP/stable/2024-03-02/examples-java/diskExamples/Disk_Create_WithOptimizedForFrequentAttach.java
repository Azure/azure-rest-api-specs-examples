
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Create_WithOptimizedForFrequentAttach.json
     */
    /**
     * Sample code: Create a managed disk with optimizedForFrequentAttach.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAManagedDiskWithOptimizedForFrequentAttach(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US")
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY)).withDiskSizeGB(200)
                .withOptimizedForFrequentAttach(true),
            com.azure.core.util.Context.NONE);
    }
}
