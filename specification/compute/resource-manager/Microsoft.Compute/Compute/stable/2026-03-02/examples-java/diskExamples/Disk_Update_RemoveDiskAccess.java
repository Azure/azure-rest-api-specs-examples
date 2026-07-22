
import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.NetworkAccessPolicy;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Update_RemoveDiskAccess.json
     */
    /**
     * Sample code: update managed disk to remove disk access resource association.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateManagedDiskToRemoveDiskAccessResourceAssociation(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withNetworkAccessPolicy(NetworkAccessPolicy.ALLOW_ALL), com.azure.core.util.Context.NONE);
    }
}
