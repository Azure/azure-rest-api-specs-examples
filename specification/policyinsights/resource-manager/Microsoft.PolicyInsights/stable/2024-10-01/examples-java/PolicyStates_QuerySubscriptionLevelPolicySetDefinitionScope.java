
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/**
 * Samples for PolicyStates ListQueryResultsForPolicySetDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyStates_QuerySubscriptionLevelPolicySetDefinitionScope.json
     */
    /**
     * Sample code: Query latest at subscription level policy set definition scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtSubscriptionLevelPolicySetDefinitionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForPolicySetDefinition(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "3e3807c1-65c9-49e0-a406-82d8ae3e338c", null, null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
