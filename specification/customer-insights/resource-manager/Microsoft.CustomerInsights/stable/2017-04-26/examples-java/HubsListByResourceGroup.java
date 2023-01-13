/** Samples for Hubs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsListByResourceGroup.json
     */
    /**
     * Sample code: Hubs_ListByResourceGroup.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void hubsListByResourceGroup(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.hubs().listByResourceGroup("TestHubRG", com.azure.core.util.Context.NONE);
    }
}
