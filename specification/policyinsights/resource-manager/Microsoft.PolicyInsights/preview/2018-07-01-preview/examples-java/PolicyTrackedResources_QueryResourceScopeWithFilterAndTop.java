import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyTrackedResourcesResourceType;

/** Samples for PolicyTrackedResources ListQueryResultsForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceScopeWithFilterAndTop.json
     */
    /**
     * Sample code: Query at resource scope using query parameters.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceScopeUsingQueryParameters(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyTrackedResources()
            .listQueryResultsForResource(
                "subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource",
                PolicyTrackedResourcesResourceType.DEFAULT,
                1,
                "PolicyAssignmentId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment'"
                    + " AND TrackedResourceId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1'",
                Context.NONE);
    }
}
