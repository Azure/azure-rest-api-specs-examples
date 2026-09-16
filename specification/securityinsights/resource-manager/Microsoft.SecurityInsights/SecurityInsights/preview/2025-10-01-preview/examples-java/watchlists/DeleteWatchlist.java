
/**
 * Samples for Watchlists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/watchlists/DeleteWatchlist.json
     */
    /**
     * Sample code: Delete a watchlist.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAWatchlist(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().delete("myRg", "myWorkspace", "highValueAsset", com.azure.core.util.Context.NONE);
    }
}
