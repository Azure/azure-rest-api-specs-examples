import com.azure.core.util.Context;

/** Samples for Watchlists Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/DeleteWatchlist.json
     */
    /**
     * Sample code: Delete a watchlist.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAWatchlist(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().deleteWithResponse("myRg", "myWorkspace", "highValueAsset", Context.NONE);
    }
}
