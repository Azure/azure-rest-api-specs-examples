/** Samples for Relationships ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsListByHub.json
     */
    /**
     * Sample code: Relationships_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipsListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.relationships().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
