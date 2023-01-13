/** Samples for Links ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksListByHub.json
     */
    /**
     * Sample code: Links_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void linksListByHub(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.links().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
