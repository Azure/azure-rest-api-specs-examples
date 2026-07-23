
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.NetworkAccessPolicy;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_WithDiskAccess.json
     */
    /**
     * Sample code: create a managed disk and associate with disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAManagedDiskAndAssociateWithDiskAccessResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk", new DiskInner()
            .withLocation("West US").withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY))
            .withDiskSizeGB(200).withNetworkAccessPolicy(NetworkAccessPolicy.ALLOW_PRIVATE).withDiskAccessId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/{existing-diskAccess-name}"),
            com.azure.core.util.Context.NONE);
    }
}
