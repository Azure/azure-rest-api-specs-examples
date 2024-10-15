
import com.azure.resourcemanager.policyinsights.models.PolicyStatesSummaryResourceType;

/**
 * Samples for PolicyStates SummarizeForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_SummarizeSubscriptionScopeForPolicyGroup.json
     */
    /**
     * Sample code: Summarize at subscription scope for a policy definition group.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtSubscriptionScopeForAPolicyDefinitionGroup(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().summarizeForSubscriptionWithResponse(PolicyStatesSummaryResourceType.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", 1, null, null, "'group1' IN PolicyDefinitionGroupNames",
            com.azure.core.util.Context.NONE);
    }
}
