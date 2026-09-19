
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/SearchListServicesBySubscription.json
     */
    /**
     * Sample code: SearchListServicesBySubscription.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchListServicesBySubscription(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().list(com.azure.core.util.Context.NONE);
    }
}
