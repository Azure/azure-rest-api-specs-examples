import com.azure.resourcemanager.securityinsights.models.SourceType;

/** Samples for Watchlists CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/watchlists/CreateWatchlistAndWatchlistItemsFromSasUri.json
     */
    /**
     * Sample code: Create or update a watchlist and bulk creates watchlist items from SAL URI.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createOrUpdateAWatchlistAndBulkCreatesWatchlistItemsFromSALURI(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .watchlists()
            .define("highValueAsset")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
            .withDisplayName("High Value Assets Watchlist")
            .withProvider("Microsoft")
            .withSourceType(SourceType.REMOTE_STORAGE)
            .withDescription("Watchlist from a large CSV file under Blob storage")
            .withNumberOfLinesToSkip(1)
            .withSasUri(
                "https://storagesample.blob.core.windows.net/sample-contaier/sampleBlob.csv?sp=r&st=2021-09-24T01:15:52Z&se=2021-10-01T09:15:52Z&spr=https&sv=2020-08-04&sr=b&sig=HRRRMc43ZJz634eBc402X%2FFPxam5sZVPSkLOY14baEd%4Z")
            .withItemsSearchKey("header1")
            .create();
    }
}
