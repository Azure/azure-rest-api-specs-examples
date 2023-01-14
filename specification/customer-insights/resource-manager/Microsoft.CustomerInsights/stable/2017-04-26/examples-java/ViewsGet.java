/** Samples for Views Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ViewsGet.json
     */
    /**
     * Sample code: Views_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void viewsGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.views().getWithResponse("TestHubRG", "sdkTestHub", "testView", "*", com.azure.core.util.Context.NONE);
    }
}
