
/**
 * Samples for WatchlistItems List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * watchlists/GetWatchlistItems.json
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
