import com.azure.core.util.Context;

/** Samples for PolicyTrackedResources ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QuerySubscriptionScope.json
     */
    /**
     * Sample code: Query at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyTrackedResources().listQueryResultsForSubscription(null, null, Context.NONE);
    }
}
