
import com.azure.resourcemanager.policyinsights.models.PolicyStatesSummaryResourceType;

/**
 * Samples for PolicyStates SummarizeForPolicyDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_SummarizeSubscriptionLevelPolicyDefinitionScope.json
     */
    /**
     * Sample code: Summarize at policy definition scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        summarizeAtPolicyDefinitionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().summarizeForPolicyDefinitionWithResponse(PolicyStatesSummaryResourceType.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
