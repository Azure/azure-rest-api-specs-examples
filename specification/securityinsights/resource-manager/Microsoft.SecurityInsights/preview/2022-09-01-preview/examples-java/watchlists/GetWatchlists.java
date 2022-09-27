import com.azure.core.util.Context;

/** Samples for Watchlists List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlists.json
     */
    /**
     * Sample code: Get all watchlists.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllWatchlists(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.watchlists().list("myRg", "myWorkspace", null, Context.NONE);
    }
}
