import com.azure.core.util.Context;

/** Samples for PolicyStates TriggerSubscriptionEvaluation. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_TriggerSubscriptionEvaluation.json
     */
    /**
     * Sample code: Trigger evaluations for all resources in a subscription.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void triggerEvaluationsForAllResourcesInASubscription(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().triggerSubscriptionEvaluation("fffedd8f-ffff-fffd-fffd-fffed2f84852", Context.NONE);
    }
}
