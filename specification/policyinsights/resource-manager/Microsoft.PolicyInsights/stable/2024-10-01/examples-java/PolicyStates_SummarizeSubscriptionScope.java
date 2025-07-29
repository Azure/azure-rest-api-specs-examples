
import com.azure.resourcemanager.policyinsights.models.PolicyStatesSummaryResourceType;

/**
 * Samples for PolicyStates SummarizeForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyStates_SummarizeSubscriptionScope.json
     */
    /**
     * Sample code: Summarize at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        summarizeAtSubscriptionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().summarizeForSubscriptionWithResponse(PolicyStatesSummaryResourceType.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", 5, null, null, null, com.azure.core.util.Context.NONE);
    }
}
