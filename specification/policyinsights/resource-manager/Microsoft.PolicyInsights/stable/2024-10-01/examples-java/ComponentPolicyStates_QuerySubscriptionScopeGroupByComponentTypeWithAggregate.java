
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * ComponentPolicyStates_QuerySubscriptionScopeGroupByComponentTypeWithAggregate.json
     */
    /**
     * Sample code: Query latest component policy compliance state count grouped by component type at subscription scope
     * filtered by given assignment.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryLatestComponentPolicyComplianceStateCountGroupedByComponentTypeAtSubscriptionScopeFilteredByGivenAssignment(
            com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForSubscriptionWithResponse(
            "e78961ba-36fe-4739-9212-e3031b4c8db7", ComponentPolicyStatesResource.LATEST, null, null, null, null, null,
            "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
            "groupby((componentType,complianceState),aggregate($count as count))", com.azure.core.util.Context.NONE);
    }
}
