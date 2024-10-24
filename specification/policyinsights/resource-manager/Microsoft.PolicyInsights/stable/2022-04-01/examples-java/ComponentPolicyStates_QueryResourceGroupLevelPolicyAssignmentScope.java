
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForResourceGroupLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QueryResourceGroupLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Query latest at resource group level policy assignment scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtResourceGroupLevelPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForResourceGroupLevelPolicyAssignmentWithResponse(
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment",
            ComponentPolicyStatesResource.LATEST, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
