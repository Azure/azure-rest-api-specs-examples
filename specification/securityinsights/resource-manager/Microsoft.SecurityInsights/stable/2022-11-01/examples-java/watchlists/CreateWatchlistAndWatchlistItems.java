
import com.azure.resourcemanager.securityinsights.models.Source;

/**
 * Samples for Watchlists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/watchlists/
     * CreateWatchlistAndWatchlistItems.json
     */
    /**
     * Sample code: Create or update a watchlist and bulk creates watchlist items.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createOrUpdateAWatchlistAndBulkCreatesWatchlistItems(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().define("highValueAsset").withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"").withDisplayName("High Value Assets Watchlist")
            .withProvider("Microsoft").withSource(Source.LOCAL_FILE).withDescription("Watchlist from CSV content")
            .withNumberOfLinesToSkip(1).withRawContent("This line will be skipped\nheader1,header2\nvalue1,value2")
            .withItemsSearchKey("header1").withContentType("text/csv").create();
    }
}
