/** Samples for Interactions ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsListByHub.json
     */
    /**
     * Sample code: Interactions_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void interactionsListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.interactions().listByHub("TestHubRG", "sdkTestHub", null, com.azure.core.util.Context.NONE);
    }
}
