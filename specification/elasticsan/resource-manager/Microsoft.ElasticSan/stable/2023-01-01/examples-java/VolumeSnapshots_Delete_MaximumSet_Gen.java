/** Samples for VolumeSnapshots Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeSnapshotsDeleteMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumeSnapshots()
            .delete(
                "resourcegroupname",
                "elasticsanname",
                "volumegroupname",
                "snapshotname",
                com.azure.core.util.Context.NONE);
    }
}
