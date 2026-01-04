
/**
 * Samples for ResourceProvider ListCustomHostnameSites.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListCustomHostNameSites.
     * json
     */
    /**
     * Sample code: Get custom hostnames under subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCustomHostnamesUnderSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceProviders().listCustomHostnameSites(null,
            com.azure.core.util.Context.NONE);
    }
}
