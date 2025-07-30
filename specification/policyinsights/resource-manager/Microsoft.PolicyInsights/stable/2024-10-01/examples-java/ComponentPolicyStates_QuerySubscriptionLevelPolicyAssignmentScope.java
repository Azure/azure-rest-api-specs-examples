
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForSubscriptionLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * ComponentPolicyStates_QuerySubscriptionLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Query latest at subscription level policy assignment scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtSubscriptionLevelPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForSubscriptionLevelPolicyAssignmentWithResponse(
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "ec8f9645-8ecb-4abb-9c0b-5292f19d4003",
            ComponentPolicyStatesResource.LATEST, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
