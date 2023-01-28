import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_Create_ByCopyingASnapshot.json
     */
    /**
     * Sample code: Create a managed disk by copying a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskByCopyingASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myDisk",
                new DiskInner()
                    .withLocation("West US")
                    .withCreationData(
                        new CreationData()
                            .withCreateOption(DiskCreateOption.COPY)
                            .withSourceResourceId(
                                "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")),
                com.azure.core.util.Context.NONE);
    }
}
