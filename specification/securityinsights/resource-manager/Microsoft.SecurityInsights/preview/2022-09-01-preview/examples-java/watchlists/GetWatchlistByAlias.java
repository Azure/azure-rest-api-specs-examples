
/**
 * Samples for Watchlists Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * watchlists/GetWatchlistByAlias.json
     */
    /**
     * Sample code: Get a watchlist.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAWatchlist(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().getWithResponse("myRg", "myWorkspace", "highValueAsset", com.azure.core.util.Context.NONE);
    }
}
