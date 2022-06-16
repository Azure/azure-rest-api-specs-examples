import com.azure.core.util.Context;

/** Samples for Remediations ListDeploymentsAtResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListDeploymentsResourceGroupScope.json
     */
    /**
     * Sample code: List deployments for a remediation at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listDeploymentsForARemediationAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().listDeploymentsAtResourceGroup("myResourceGroup", "myRemediation", null, Context.NONE);
    }
}
