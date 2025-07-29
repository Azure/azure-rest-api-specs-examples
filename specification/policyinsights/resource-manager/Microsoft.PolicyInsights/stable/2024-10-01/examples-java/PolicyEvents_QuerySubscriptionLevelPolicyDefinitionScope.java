
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForPolicyDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyEvents_QuerySubscriptionLevelPolicyDefinitionScope.json
     */
    /**
     * Sample code: Query at subscription level policy definition scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionLevelPolicyDefinitionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForPolicyDefinition(PolicyEventsResourceType.DEFAULT,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1", null, null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
