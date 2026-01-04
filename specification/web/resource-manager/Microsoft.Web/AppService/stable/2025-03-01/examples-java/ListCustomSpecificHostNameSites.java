
/**
 * Samples for ResourceProvider ListCustomHostnameSites.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListCustomSpecificHostNameSites.json
     */
    /**
     * Sample code: Get specific custom hostname under subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getSpecificCustomHostnameUnderSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceProviders().listCustomHostnameSites("www.example.com",
            com.azure.core.util.Context.NONE);
    }
}
