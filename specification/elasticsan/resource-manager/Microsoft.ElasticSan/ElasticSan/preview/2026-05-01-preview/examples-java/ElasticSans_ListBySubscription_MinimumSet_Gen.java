
/**
 * Samples for ElasticSans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        elasticSansListBySubscriptionMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().list(com.azure.core.util.Context.NONE);
    }
}
