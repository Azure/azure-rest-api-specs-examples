Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Watchlists Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/watchlists/DeleteWatchlist.json
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
```
