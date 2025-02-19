
/**
 * Samples for Watchlists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/watchlists/
     * DeleteWatchlist.json
     */
    /**
     * Sample code: Delete a watchlist.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAWatchlist(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().deleteWithResponse("myRg", "myWorkspace", "highValueAsset",
            com.azure.core.util.Context.NONE);
    }
}
