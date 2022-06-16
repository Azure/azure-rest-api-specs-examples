import com.azure.core.util.Context;

/** Samples for PolicyEvents ListQueryResultsForSubscriptionLevelPolicyAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QuerySubscriptionLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Query at subscription level policy assignment scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionLevelPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForSubscriptionLevelPolicyAssignment(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "ec8f9645-8ecb-4abb-9c0b-5292f19d4003",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
