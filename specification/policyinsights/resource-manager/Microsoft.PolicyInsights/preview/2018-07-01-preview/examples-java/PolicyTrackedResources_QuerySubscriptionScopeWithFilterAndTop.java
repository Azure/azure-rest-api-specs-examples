import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyTrackedResourcesResourceType;

/** Samples for PolicyTrackedResources ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QuerySubscriptionScopeWithFilterAndTop.json
     */
    /**
     * Sample code: Query at subscription scope using query parameters.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionScopeUsingQueryParameters(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyTrackedResources()
            .listQueryResultsForSubscription(
                PolicyTrackedResourcesResourceType.DEFAULT,
                1,
                "PolicyAssignmentId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment'"
                    + " AND TrackedResourceId eq"
                    + " '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/exampleTrackedResourceName'",
                Context.NONE);
    }
}
