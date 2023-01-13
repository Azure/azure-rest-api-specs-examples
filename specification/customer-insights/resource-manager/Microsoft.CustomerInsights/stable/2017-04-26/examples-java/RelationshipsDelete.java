/** Samples for Relationships Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsDelete.json
     */
    /**
     * Sample code: Relationships_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipsDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.relationships().delete("TestHubRG", "sdkTestHub", "SomeRelationship", com.azure.core.util.Context.NONE);
    }
}
