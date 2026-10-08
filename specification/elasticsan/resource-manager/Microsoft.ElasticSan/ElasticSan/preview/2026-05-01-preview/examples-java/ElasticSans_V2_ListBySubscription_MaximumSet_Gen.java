
/**
 * Samples for ElasticSans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V2_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        elasticSansV2ListBySubscriptionMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().list(com.azure.core.util.Context.NONE);
    }
}
