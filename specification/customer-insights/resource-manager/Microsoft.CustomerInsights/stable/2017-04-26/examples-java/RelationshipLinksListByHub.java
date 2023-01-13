/** Samples for RelationshipLinks ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksListByHub.json
     */
    /**
     * Sample code: RelationshipLinks_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipLinksListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.relationshipLinks().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
