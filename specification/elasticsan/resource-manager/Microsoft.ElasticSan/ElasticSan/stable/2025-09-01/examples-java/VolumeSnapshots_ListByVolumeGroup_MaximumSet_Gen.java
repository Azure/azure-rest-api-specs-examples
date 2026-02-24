
/**
 * Samples for VolumeSnapshots ListByVolumeGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeSnapshotsListByVolumeGroupMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeSnapshots().listByVolumeGroup("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumeName eq <volume name>", com.azure.core.util.Context.NONE);
    }
}
