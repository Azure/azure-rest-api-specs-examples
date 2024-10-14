
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QueryResourceGroupScope.json
     */
    /**
     * Sample code: Query latest component policy states at resource group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestComponentPolicyStatesAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForResourceGroupWithResponse(
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", ComponentPolicyStatesResource.LATEST, null, null,
            null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
