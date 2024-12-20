
/**
 * Samples for Remediations DeleteAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/
     * Remediations_DeleteSubscriptionScope.json
     */
    /**
     * Sample code: Delete remediation at subscription scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        deleteRemediationAtSubscriptionScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().deleteAtSubscriptionWithResponse("storageRemediation", com.azure.core.util.Context.NONE);
    }
}
