
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QueryResourceScopeGroupByComponentTypeWithAggregate.json
     */
    /**
     * Sample code: Query latest component policy compliance state count grouped by component type at resource scope
     * filtered by given assignment.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryLatestComponentPolicyComplianceStateCountGroupedByComponentTypeAtResourceScopeFilteredByGivenAssignment(
            com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForResourceWithResponse(
            "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
            ComponentPolicyStatesResource.LATEST, null, null, null, null, null,
            "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
            "groupby((componentType,complianceState),aggregate($count as count))", null,
            com.azure.core.util.Context.NONE);
    }
}
