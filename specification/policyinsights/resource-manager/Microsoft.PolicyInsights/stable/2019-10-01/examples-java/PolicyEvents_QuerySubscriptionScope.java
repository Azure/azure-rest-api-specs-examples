
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyEvents_QuerySubscriptionScope.json
     */
    /**
     * Sample code: Query at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryAtSubscriptionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForSubscription(PolicyEventsResourceType.DEFAULT,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
