
/**
 * Samples for WatchlistItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/watchlists/
     * DeleteWatchlistItem.json
     */
    /**
     * Sample code: Delete a watchlist item.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAWatchlistItem(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlistItems().deleteWithResponse("myRg", "myWorkspace", "highValueAsset",
            "4008512e-1d30-48b2-9ee2-d3612ed9d3ea", com.azure.core.util.Context.NONE);
    }
}
