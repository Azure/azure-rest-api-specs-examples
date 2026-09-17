
/**
 * Samples for WatchlistItems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/watchlists/GetWatchlistItems.json
     */
    /**
     * Sample code: Get all watchlist Items.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllWatchlistItems(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlistItems().list("myRg", "myWorkspace", "highValueAsset", null, com.azure.core.util.Context.NONE);
    }
}
