
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForResourceGroupLevelPolicyAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyEvents_QueryResourceGroupLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Query at resource group level policy assignment scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceGroupLevelPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForResourceGroupLevelPolicyAssignment(PolicyEventsResourceType.DEFAULT,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", null, null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
