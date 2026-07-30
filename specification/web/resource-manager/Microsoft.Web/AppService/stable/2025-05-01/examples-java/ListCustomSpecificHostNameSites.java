
/**
 * Samples for ResourceProviders ListCustomHostnameSites.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListCustomSpecificHostNameSites.json
     */
    /**
     * Sample code: Get specific custom hostname under subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getSpecificCustomHostnameUnderSubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceProviders().listCustomHostnameSites("www.example.com",
            com.azure.core.util.Context.NONE);
    }
}
