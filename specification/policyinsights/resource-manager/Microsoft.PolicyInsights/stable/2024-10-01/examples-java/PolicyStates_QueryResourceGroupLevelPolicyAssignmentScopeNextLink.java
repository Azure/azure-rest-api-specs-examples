
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/**
 * Samples for PolicyStates ListQueryResultsForResourceGroupLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyStates_QueryResourceGroupLevelPolicyAssignmentScopeNextLink.json
     */
    /**
     * Sample code: Query latest at resource group level policy assignment scope with next link.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtResourceGroupLevelPolicyAssignmentScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForResourceGroupLevelPolicyAssignment(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", null, null, null, null,
            null, null, null, "WpmWfBSvPhkAK6QD", com.azure.core.util.Context.NONE);
    }
}
