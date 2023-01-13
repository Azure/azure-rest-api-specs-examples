/** Samples for Hubs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsGet.json
     */
    /**
     * Sample code: Hubs_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void hubsGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.hubs().getByResourceGroupWithResponse("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
