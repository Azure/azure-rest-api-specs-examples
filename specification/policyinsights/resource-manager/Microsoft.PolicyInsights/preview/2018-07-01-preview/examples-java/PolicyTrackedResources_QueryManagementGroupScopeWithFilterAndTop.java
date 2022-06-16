import com.azure.core.util.Context;

/** Samples for PolicyTrackedResources ListQueryResultsForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScopeWithFilterAndTop.json
     */
    /**
     * Sample code: Query at management group scope using query parameters.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtManagementGroupScopeUsingQueryParameters(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyTrackedResources()
            .listQueryResultsForManagementGroup(
                "myManagementGroup",
                1,
                "PolicyAssignmentId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment'"
                    + " AND TrackedResourceId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/exampleTrackedResourceName'",
                Context.NONE);
    }
}
