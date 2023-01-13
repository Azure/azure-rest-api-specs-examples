/** Samples for Views ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ViewsListByHub.json
     */
    /**
     * Sample code: Views_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void viewsListByHub(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.views().listByHub("TestHubRG", "sdkTestHub", "*", com.azure.core.util.Context.NONE);
    }
}
