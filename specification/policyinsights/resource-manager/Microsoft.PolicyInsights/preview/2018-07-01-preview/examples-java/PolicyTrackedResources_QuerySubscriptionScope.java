import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyTrackedResourcesResourceType;

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
        manager
            .policyTrackedResources()
            .listQueryResultsForSubscription(PolicyTrackedResourcesResourceType.DEFAULT, null, null, Context.NONE);
    }
}
