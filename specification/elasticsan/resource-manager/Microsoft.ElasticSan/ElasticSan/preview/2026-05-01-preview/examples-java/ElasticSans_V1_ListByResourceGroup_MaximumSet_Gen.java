
/**
 * Samples for ElasticSans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V1_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V1_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        elasticSansV1ListByResourceGroupMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().listByResourceGroup("resourcegroupname", com.azure.core.util.Context.NONE);
    }
}
