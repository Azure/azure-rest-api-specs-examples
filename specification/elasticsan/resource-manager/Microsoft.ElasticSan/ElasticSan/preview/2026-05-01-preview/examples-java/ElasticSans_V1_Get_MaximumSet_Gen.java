
/**
 * Samples for ElasticSans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V1_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V1_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansV1GetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().getByResourceGroupWithResponse("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
