
import com.azure.resourcemanager.compute.models.Architecture;
import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.SupportedCapabilities;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Update_ToAddArchitecture.json
     */
    /**
     * Sample code: update a managed disk to add architecture.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateAManagedDiskToAddArchitecture(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks()
            .update("myResourceGroup", "myDisk",
                new DiskUpdate()
                    .withSupportedCapabilities(new SupportedCapabilities().withArchitecture(Architecture.ARM64)),
                com.azure.core.util.Context.NONE);
    }
}
