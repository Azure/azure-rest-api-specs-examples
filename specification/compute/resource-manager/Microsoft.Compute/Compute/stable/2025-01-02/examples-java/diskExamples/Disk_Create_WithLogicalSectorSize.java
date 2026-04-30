
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSku;
import com.azure.resourcemanager.compute.models.DiskStorageAccountTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithLogicalSectorSize.json
     */
    /**
     * Sample code: create an ultra managed disk with logicalSectorSize 512E.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAnUltraManagedDiskWithLogicalSectorSize512E(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US")
                .withSku(new DiskSku().withName(DiskStorageAccountTypes.ULTRA_SSD_LRS))
                .withCreationData(
                    new CreationData().withCreateOption(DiskCreateOption.EMPTY).withLogicalSectorSize(512))
                .withDiskSizeGB(200),
            com.azure.core.util.Context.NONE);
    }
}
