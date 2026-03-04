
/**
 * Samples for Organizations GetElasticToAzureSubscriptionMapping.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Organizations_GetElasticToAzureSubscriptionMapping.json
     */
    /**
     * Sample code: Organizations_GetElasticToAzureSubscriptionMapping.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void
        organizationsGetElasticToAzureSubscriptionMapping(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.organizations().getElasticToAzureSubscriptionMappingWithResponse(com.azure.core.util.Context.NONE);
    }
}
