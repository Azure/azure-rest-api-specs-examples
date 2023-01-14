/** Samples for Links Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksGet.json
     */
    /**
     * Sample code: Links_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void linksGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.links().getWithResponse("TestHubRG", "sdkTestHub", "linkTest4806", com.azure.core.util.Context.NONE);
    }
}
