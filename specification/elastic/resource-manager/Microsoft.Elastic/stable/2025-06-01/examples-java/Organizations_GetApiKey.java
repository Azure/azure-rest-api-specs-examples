
/**
 * Samples for Organizations GetApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/Organizations_GetApiKey.json
     */
    /**
     * Sample code: Organizations_GetApiKey.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void organizationsGetApiKey(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.organizations().getApiKeyWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
