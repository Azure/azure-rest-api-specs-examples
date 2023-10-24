/** Samples for DiscoverySolution List. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/ListDiscoverySolutionsAtSubscriptionScope.json
     */
    /**
     * Sample code: List DiscoverySolutions at subscription scope.
     *
     * @param manager Entry point to SelfHelpManager.
     */
    public static void listDiscoverySolutionsAtSubscriptionScope(
        com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager
            .discoverySolutions()
            .list(
                "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6",
                "ProblemClassificationId eq 'SampleProblemClassificationId1'",
                null,
                com.azure.core.util.Context.NONE);
    }
}
