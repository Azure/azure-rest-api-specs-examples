import com.azure.core.util.Context;

/** Samples for WatchlistItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlistItemById.json
     */
    /**
     * Sample code: Get a watchlist item.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAWatchlistItem(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .watchlistItems()
            .getWithResponse(
                "myRg", "myWorkspace", "highValueAsset", "3f8901fe-63d9-4875-9ad5-9fb3b8105797", Context.NONE);
    }
}
