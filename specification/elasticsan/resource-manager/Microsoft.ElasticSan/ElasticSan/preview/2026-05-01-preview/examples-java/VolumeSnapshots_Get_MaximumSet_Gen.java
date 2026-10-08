
/**
 * Samples for VolumeSnapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeSnapshots_Get_MaximumSet_Gen.json
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
