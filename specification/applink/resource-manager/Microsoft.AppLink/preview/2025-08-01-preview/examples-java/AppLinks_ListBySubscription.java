
/**
 * Samples for AppLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinks_ListBySubscription.json
     */
    /**
     * Sample code: AppLinks_ListBySubscription.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinksListBySubscription(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinks().list(com.azure.core.util.Context.NONE);
    }
}
