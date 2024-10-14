
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QuerySubscriptionScope.json
     */
    /**
     * Sample code: Query latest component policy states at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestComponentPolicyStatesAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForSubscriptionWithResponse(
            "fff10b27-fff3-fff5-fff8-fffbe01e86a5", ComponentPolicyStatesResource.LATEST, null, null, null, null, null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
