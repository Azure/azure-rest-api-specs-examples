
/**
 * Samples for VolumeSnapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * VolumeSnapshots_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeSnapshotsGetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeSnapshots().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            "snapshotname", com.azure.core.util.Context.NONE);
    }
}
