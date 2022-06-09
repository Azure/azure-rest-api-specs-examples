```java
import com.azure.core.util.Context;

/** Samples for WatchlistItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/watchlists/DeleteWatchlistItem.json
     */
    /**
     * Sample code: Delete a watchlist Item.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAWatchlistItem(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .watchlistItems()
            .deleteWithResponse(
                "myRg", "myWorkspace", "highValueAsset", "4008512e-1d30-48b2-9ee2-d3612ed9d3ea", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
