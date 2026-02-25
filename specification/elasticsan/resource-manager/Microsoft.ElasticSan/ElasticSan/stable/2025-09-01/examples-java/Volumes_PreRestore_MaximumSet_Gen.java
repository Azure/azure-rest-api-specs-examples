
import com.azure.resourcemanager.elasticsan.models.DiskSnapshotList;
import java.util.Arrays;

/**
 * Samples for Volumes PreRestore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Volumes_PreRestore_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_PreRestore_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsPreRestoreMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().preRestore("resourcegroupname", "elasticsanname", "volumegroupname",
            new DiskSnapshotList().withDiskSnapshotIds(Arrays.asList(
                "/subscriptions/{subscriptionid}/resourceGroups/{resourcegroupname}/providers/Microsoft.Compute/snapshots/disksnapshot1")),
            com.azure.core.util.Context.NONE);
    }
}
