
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DataAccessAuthMode;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_WithDataAccessAuthMode.json
     */
    /**
     * Sample code: create a managed disk with dataAccessAuthMode.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAManagedDiskWithDataAccessAuthMode(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US")
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY)).withDiskSizeGB(200)
                .withDataAccessAuthMode(DataAccessAuthMode.AZURE_ACTIVE_DIRECTORY),
            com.azure.core.util.Context.NONE);
    }
}
