import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesSummaryResourceType;

/** Samples for PolicyStates SummarizeForSubscriptionLevelPolicyAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Summarize at policy assignment scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForSubscriptionLevelPolicyAssignmentWithResponse(
                PolicyStatesSummaryResourceType.LATEST,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "ec8f9645-8ecb-4abb-9c0b-5292f19d4003",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
