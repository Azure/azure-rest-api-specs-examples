
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForResourceGroupLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyEvents_QueryResourceGroupLevelPolicyAssignmentScopeNextLink.json
     */
    /**
     * Sample code: Query at resource group level policy assignment scope with next link.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceGroupLevelPolicyAssignmentScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForResourceGroupLevelPolicyAssignment(PolicyEventsResourceType.DEFAULT,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", null, null, null, null,
            null, null, null, "WpmWfBSvPhkAK6QD", com.azure.core.util.Context.NONE);
    }
}
