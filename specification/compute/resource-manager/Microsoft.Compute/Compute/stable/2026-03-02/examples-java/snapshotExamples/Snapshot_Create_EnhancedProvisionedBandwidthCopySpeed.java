
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.ProvisionedBandwidthCopyOption;

/**
 * Samples for Snapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_Create_EnhancedProvisionedBandwidthCopySpeed.json
     */
    /**
     * Sample code: Create a snapshot from an existing snapshot in the same or a different subscription in a different
     * region with quicker copy speed.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscriptionInADifferentRegionWithQuickerCopySpeed(
            com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().createOrUpdate("myResourceGroup", "mySnapshot2", new SnapshotInner()
            .withLocation("West US")
            .withCreationData(new CreationData().withCreateOption(DiskCreateOption.COPY_START).withSourceResourceId(
                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1")
                .withProvisionedBandwidthCopySpeed(ProvisionedBandwidthCopyOption.ENHANCED)),
            com.azure.core.util.Context.NONE);
    }
}
