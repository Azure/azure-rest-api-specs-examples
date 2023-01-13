/** Samples for Views Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ViewsDelete.json
     */
    /**
     * Sample code: Views_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void viewsDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .views()
            .deleteWithResponse("TestHubRG", "sdkTestHub", "testView", "*", com.azure.core.util.Context.NONE);
    }
}
