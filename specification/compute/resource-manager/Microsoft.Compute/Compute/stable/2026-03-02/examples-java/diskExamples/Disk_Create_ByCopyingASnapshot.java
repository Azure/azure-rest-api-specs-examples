
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_ByCopyingASnapshot.json
     */
    /**
     * Sample code: create a managed disk by copying a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAManagedDiskByCopyingASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk", new DiskInner()
            .withLocation("West US")
            .withCreationData(new CreationData().withCreateOption(DiskCreateOption.COPY).withSourceResourceId(
                "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")),
            com.azure.core.util.Context.NONE);
    }
}
