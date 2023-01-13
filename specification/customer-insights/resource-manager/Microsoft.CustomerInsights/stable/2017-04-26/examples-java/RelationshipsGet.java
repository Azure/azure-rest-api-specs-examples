/** Samples for Relationships Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsGet.json
     */
    /**
     * Sample code: Relationships_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipsGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .relationships()
            .getWithResponse("TestHubRG", "sdkTestHub", "SomeRelationship", com.azure.core.util.Context.NONE);
    }
}
