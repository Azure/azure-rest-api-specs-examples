
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Update_DisableBursting.json
     */
    /**
     * Sample code: update a managed disk to disable bursting.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateAManagedDiskToDisableBursting(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withBurstingEnabled(false), com.azure.core.util.Context.NONE);
    }
}
