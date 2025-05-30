
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForPolicyDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QuerySubscriptionLevelPolicyDefinitionScope.json
     */
    /**
     * Sample code: Query latest component policy states at subscription level policy definition scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestComponentPolicyStatesAtSubscriptionLevelPolicyDefinitionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForPolicyDefinitionWithResponse(
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1",
            ComponentPolicyStatesResource.LATEST, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
