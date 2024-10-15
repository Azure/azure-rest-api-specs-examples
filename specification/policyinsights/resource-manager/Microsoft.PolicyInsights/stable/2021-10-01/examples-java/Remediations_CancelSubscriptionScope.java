
/**
 * Samples for Remediations CancelAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/
     * Remediations_CancelSubscriptionScope.json
     */
    /**
     * Sample code: Cancel a remediation at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        cancelARemediationAtSubscriptionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().cancelAtSubscriptionWithResponse("myRemediation", com.azure.core.util.Context.NONE);
    }
}
