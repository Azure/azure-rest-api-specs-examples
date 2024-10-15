
import com.azure.resourcemanager.policyinsights.models.ComponentPolicyStatesResource;

/**
 * Samples for ComponentPolicyStates ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/
     * ComponentPolicyStates_QueryResourceScopeExpandPolicyEvaluationDetails.json
     */
    /**
     * Sample code: Query latest component policy states at resource scope and expand policyEvaluationDetails.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestComponentPolicyStatesAtResourceScopeAndExpandPolicyEvaluationDetails(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.componentPolicyStates().listQueryResultsForResourceWithResponse(
            "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster",
            ComponentPolicyStatesResource.LATEST, null, null, null, null, null,
            "componentType eq 'pod' AND componentId eq 'default/test-pod' AND componentName eq 'test-pod'", null,
            "PolicyEvaluationDetails", com.azure.core.util.Context.NONE);
    }
}
