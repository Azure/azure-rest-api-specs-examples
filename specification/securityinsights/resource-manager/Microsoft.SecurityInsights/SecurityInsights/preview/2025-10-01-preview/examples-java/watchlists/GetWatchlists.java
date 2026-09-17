
/**
 * Samples for Watchlists List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/watchlists/GetWatchlists.json
     */
    /**
     * Sample code: Get all watchlists.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllWatchlists(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().list("myRg", "myWorkspace", null, com.azure.core.util.Context.NONE);
    }
}
