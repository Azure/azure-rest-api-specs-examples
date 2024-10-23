
/**
 * Samples for VolumeSnapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * VolumeSnapshots_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeSnapshotsDeleteMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeSnapshots().delete("resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname",
            com.azure.core.util.Context.NONE);
    }
}
