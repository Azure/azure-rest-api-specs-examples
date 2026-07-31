
/**
 * Samples for ResourceProviders ListCustomHostnameSites.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListCustomHostNameSites.json
     */
    /**
     * Sample code: Get custom hostnames under subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getCustomHostnamesUnderSubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceProviders().listCustomHostnameSites(null, com.azure.core.util.Context.NONE);
    }
}
