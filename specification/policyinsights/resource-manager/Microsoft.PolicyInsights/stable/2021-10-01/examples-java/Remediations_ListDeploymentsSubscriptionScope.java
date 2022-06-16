import com.azure.core.util.Context;

/** Samples for Remediations ListDeploymentsAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListDeploymentsSubscriptionScope.json
     */
    /**
     * Sample code: List deployments for a remediation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listDeploymentsForARemediationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().listDeploymentsAtSubscription("myRemediation", null, Context.NONE);
    }
}
