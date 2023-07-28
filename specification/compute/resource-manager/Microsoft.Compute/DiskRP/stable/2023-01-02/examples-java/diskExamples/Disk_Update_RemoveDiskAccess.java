import com.azure.resourcemanager.compute.models.DiskUpdate;
import com.azure.resourcemanager.compute.models.NetworkAccessPolicy;

/** Samples for Disks Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskExamples/Disk_Update_RemoveDiskAccess.json
     */
    /**
     * Sample code: Update managed disk to remove disk access resource association.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateManagedDiskToRemoveDiskAccessResourceAssociation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .update(
                "myResourceGroup",
                "myDisk",
                new DiskUpdate().withNetworkAccessPolicy(NetworkAccessPolicy.ALLOW_ALL),
                com.azure.core.util.Context.NONE);
    }
}
