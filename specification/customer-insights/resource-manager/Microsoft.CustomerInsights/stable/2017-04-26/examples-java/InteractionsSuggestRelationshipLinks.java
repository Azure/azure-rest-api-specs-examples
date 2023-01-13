/** Samples for Interactions SuggestRelationshipLinks. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsSuggestRelationshipLinks.json
     */
    /**
     * Sample code: Interactions_SuggestRelationshipLinks.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void interactionsSuggestRelationshipLinks(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .interactions()
            .suggestRelationshipLinksWithResponse(
                "TestHubRG", "sdkTestHub", "Deposit", com.azure.core.util.Context.NONE);
    }
}
