
/**
 * Samples for ElasticSans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansDeleteMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().delete("resourcegroupname", "elasticsanname", com.azure.core.util.Context.NONE);
    }
}
