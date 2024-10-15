
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/**
 * Samples for PolicyStates ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_QueryResourceScopeExpandComponents.json
     */
    /**
     * Sample code: Query component policy compliance state at resource scope filtered by given assignment.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryComponentPolicyComplianceStateAtResourceScopeFilteredByGivenAssignment(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForResource(PolicyStatesResource.LATEST,
            "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
            null, null, null, null, null,
            "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
            null, "components($filter=ComplianceState eq 'NonCompliant' or ComplianceState eq 'Compliant')", null,
            com.azure.core.util.Context.NONE);
    }
}
