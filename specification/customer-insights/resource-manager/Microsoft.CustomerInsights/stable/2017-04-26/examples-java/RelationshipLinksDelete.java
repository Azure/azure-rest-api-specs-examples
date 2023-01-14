/** Samples for RelationshipLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksDelete.json
     */
    /**
     * Sample code: RelationshipLinks_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipLinksDelete(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.relationshipLinks().delete("TestHubRG", "sdkTestHub", "Somelink", com.azure.core.util.Context.NONE);
    }
}
