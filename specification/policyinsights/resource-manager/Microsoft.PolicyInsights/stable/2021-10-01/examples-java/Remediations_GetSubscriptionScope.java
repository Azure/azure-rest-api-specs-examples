import com.azure.core.util.Context;

/** Samples for Remediations GetAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetSubscriptionScope.json
     */
    /**
     * Sample code: Get remediation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getRemediationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().getAtSubscriptionWithResponse("storageRemediation", Context.NONE);
    }
}
