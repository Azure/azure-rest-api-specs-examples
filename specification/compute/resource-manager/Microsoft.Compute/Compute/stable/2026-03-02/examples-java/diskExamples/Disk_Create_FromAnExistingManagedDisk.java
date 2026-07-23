
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_FromAnExistingManagedDisk.json
     */
    /**
     * Sample code: create a managed disk from an existing managed disk in the same or different subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAManagedDiskFromAnExistingManagedDiskInTheSameOrDifferentSubscription(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk2", new DiskInner()
            .withLocation("West US")
            .withCreationData(new CreationData().withCreateOption(DiskCreateOption.COPY).withSourceResourceId(
                "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1")),
            com.azure.core.util.Context.NONE);
    }
}
