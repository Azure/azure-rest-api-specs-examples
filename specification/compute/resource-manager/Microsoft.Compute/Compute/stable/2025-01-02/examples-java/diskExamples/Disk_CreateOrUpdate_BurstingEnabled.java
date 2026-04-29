
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_CreateOrUpdate_BurstingEnabled.json
     */
    /**
     * Sample code: create or update a bursting enabled managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createOrUpdateABurstingEnabledManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withDiskSizeGB(1024).withBurstingEnabled(true), com.azure.core.util.Context.NONE);
    }
}
