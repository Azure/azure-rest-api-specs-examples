
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Update_DisableOptimizedForFrequentAttach.json
     */
    /**
     * Sample code: update a managed disk to disable optimizedForFrequentAttach.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateAManagedDiskToDisableOptimizedForFrequentAttach(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withOptimizedForFrequentAttach(false), com.azure.core.util.Context.NONE);
    }
}
