/** Samples for RelationshipLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksGet.json
     */
    /**
     * Sample code: RelationshipLinks_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void relationshipLinksGet(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .relationshipLinks()
            .getWithResponse("TestHubRG", "sdkTestHub", "Somelink", com.azure.core.util.Context.NONE);
    }
}
