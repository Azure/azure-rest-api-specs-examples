import com.azure.core.util.Context;

/** Samples for Remediations ListDeploymentsAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListDeploymentsManagementGroupScope.json
     */
    /**
     * Sample code: List deployments for a remediation at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listDeploymentsForARemediationAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().listDeploymentsAtManagementGroup("financeMg", "myRemediation", null, Context.NONE);
    }
}
