/** Samples for VolumeSnapshots ListByVolumeGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeSnapshotsListByVolumeGroupMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumeSnapshots()
            .listByVolumeGroup(
                "resourcegroupname",
                "elasticsanname",
                "volumegroupname",
                "volumeName eq <volume name>",
                com.azure.core.util.Context.NONE);
    }
}
