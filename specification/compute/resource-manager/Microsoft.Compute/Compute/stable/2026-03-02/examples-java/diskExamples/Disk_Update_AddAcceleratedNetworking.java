
import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.SupportedCapabilities;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Update_AddAcceleratedNetworking.json
     */
    /**
     * Sample code: update a managed disk to add accelerated networking.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        updateAManagedDiskToAddAcceleratedNetworking(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withSupportedCapabilities(new SupportedCapabilities().withAcceleratedNetwork(false)),
            com.azure.core.util.Context.NONE);
    }
}
