
/**
 * Samples for VolumeSnapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2024-05-01/examples/
     * VolumeSnapshots_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeSnapshotsGetMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeSnapshots().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            "snapshotname", com.azure.core.util.Context.NONE);
    }
}
