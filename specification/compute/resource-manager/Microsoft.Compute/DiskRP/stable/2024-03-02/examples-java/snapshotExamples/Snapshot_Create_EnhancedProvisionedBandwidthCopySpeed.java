
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.ProvisionedBandwidthCopyOption;

/**
 * Samples for Snapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/
     * Snapshot_Create_EnhancedProvisionedBandwidthCopySpeed.json
     */
    /**
     * Sample code: Create a snapshot from an existing snapshot in the same or a different subscription in a different
     * region with quicker copy speed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscriptionInADifferentRegionWithQuickerCopySpeed(
            com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots()
            .createOrUpdate("myResourceGroup", "mySnapshot2", new SnapshotInner().withLocation("West US")
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.COPY_START).withSourceResourceId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1")
                    .withProvisionedBandwidthCopySpeed(ProvisionedBandwidthCopyOption.ENHANCED)),
                com.azure.core.util.Context.NONE);
    }
}
