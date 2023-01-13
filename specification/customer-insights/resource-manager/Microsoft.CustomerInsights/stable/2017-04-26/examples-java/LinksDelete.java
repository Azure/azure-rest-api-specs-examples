/** Samples for Links Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksDelete.json
     */
    /**
     * Sample code: Links_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void linksDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.links().deleteWithResponse("TestHubRG", "sdkTestHub", "linkTest4806", com.azure.core.util.Context.NONE);
    }
}
