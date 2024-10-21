
/**
 * Samples for Organizations GetElasticToAzureSubscriptionMapping.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/
     * Organizations_GetElasticToAzureSubscriptionMapping.json
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
