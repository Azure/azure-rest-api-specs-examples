
import com.azure.resourcemanager.securityinsights.models.SourceType;

/**
 * Samples for Watchlists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * watchlists/CreateWatchlist.json
     */
    /**
     * Sample code: Creates or updates a watchlist.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createsOrUpdatesAWatchlist(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().define("highValueAsset").withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"").withDisplayName("High Value Assets Watchlist")
            .withProvider("Microsoft").withSource("watchlist.csv").withSourceType(SourceType.LOCAL_FILE)
            .withDescription("Watchlist from CSV content").withItemsSearchKey("header1").create();
    }
}
