
/**
 * Samples for ElasticSans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * ElasticSans_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansGetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().getByResourceGroupWithResponse("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
