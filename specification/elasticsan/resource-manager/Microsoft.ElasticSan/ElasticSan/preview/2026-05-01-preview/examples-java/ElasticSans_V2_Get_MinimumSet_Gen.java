
/**
 * Samples for ElasticSans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V2_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansV2GetMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().getByResourceGroupWithResponse("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
