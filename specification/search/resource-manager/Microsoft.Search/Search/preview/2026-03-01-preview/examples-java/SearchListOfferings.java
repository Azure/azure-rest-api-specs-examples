
/**
 * Samples for Offerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchListOfferings.json
     */
    /**
     * Sample code: SearchListOfferings.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchListOfferings(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getOfferings().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
