
/**
 * Samples for Volumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * Volumes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesGetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname", "volumename",
            com.azure.core.util.Context.NONE);
    }
}
