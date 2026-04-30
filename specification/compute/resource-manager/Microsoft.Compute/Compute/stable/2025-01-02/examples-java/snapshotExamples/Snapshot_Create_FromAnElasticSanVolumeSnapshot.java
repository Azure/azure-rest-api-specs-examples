
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Snapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Create_FromAnElasticSanVolumeSnapshot.json
     */
    /**
     * Sample code: Create a snapshot from an elastic san volume snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createASnapshotFromAnElasticSanVolumeSnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().createOrUpdate("myResourceGroup", "mySnapshot",
            new SnapshotInner().withLocation("West US").withCreationData(
                new CreationData().withCreateOption(DiskCreateOption.COPY_FROM_SAN_SNAPSHOT).withElasticSanResourceId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ElasticSan/elasticSans/myElasticSan/volumegroups/myElasticSanVolumeGroup/snapshots/myElasticSanVolumeSnapshot")),
            com.azure.core.util.Context.NONE);
    }
}
