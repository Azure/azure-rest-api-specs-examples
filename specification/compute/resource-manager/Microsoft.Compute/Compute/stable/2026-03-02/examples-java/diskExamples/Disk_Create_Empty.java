
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_Empty.json
     */
    /**
     * Sample code: create an empty managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAnEmptyManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US")
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY)).withDiskSizeGB(200),
            com.azure.core.util.Context.NONE);
    }
}
