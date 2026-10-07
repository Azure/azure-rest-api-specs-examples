
/**
 * Samples for Volumes Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Volumes_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Create_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesCreateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().define("volumename")
            .withExistingVolumegroup("resourcegroupname", "elasticsanname", "volumegroupname").withSizeGiB(9L).create();
    }
}
