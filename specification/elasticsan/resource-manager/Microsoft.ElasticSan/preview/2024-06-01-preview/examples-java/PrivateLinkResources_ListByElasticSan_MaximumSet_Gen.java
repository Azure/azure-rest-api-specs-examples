
/**
 * Samples for PrivateLinkResources ListByElasticSan.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * PrivateLinkResources_ListByElasticSan_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByElasticSan_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void privateLinkResourcesListByElasticSanMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.privateLinkResources().listByElasticSanWithResponse("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
