import com.azure.core.util.Context;

/** Samples for PolicyStates TriggerResourceGroupEvaluation. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_TriggerResourceGroupEvaluation.json
     */
    /**
     * Sample code: Trigger evaluations for all resources in a resource group.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void triggerEvaluationsForAllResourcesInAResourceGroup(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .triggerResourceGroupEvaluation("fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", Context.NONE);
    }
}
