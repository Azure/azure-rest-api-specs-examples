
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/**
 * Samples for PolicyStates ListQueryResultsForResourceGroupLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_QueryResourceGroupLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Query latest at resource group level policy assignment scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtResourceGroupLevelPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForResourceGroupLevelPolicyAssignment(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", null, null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
