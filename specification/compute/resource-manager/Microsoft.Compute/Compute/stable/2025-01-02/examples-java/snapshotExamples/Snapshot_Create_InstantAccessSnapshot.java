
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Snapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Create_InstantAccessSnapshot.json
     */
    /**
     * Sample code: create a snapshot which can be instantly accessable.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createASnapshotWhichCanBeInstantlyAccessable(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().createOrUpdate("myResourceGroup", "mySnapshot2", new SnapshotInner()
            .withLocation("West US")
            .withCreationData(new CreationData().withCreateOption(DiskCreateOption.COPY).withSourceResourceId(
                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1")
                .withInstantAccessDurationMinutes(120L)),
            com.azure.core.util.Context.NONE);
    }
}
