
/**
 * Samples for WebApps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWebApps.json
     */
    /**
     * Sample code: List Web apps for subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWebAppsForSubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().list(com.azure.core.util.Context.NONE);
    }
}
