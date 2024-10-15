
/**
 * Samples for Remediations CancelAtResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/
     * Remediations_CancelResourceGroupScope.json
     */
    /**
     * Sample code: Cancel a remediation at resource group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        cancelARemediationAtResourceGroupScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().cancelAtResourceGroupWithResponse("myResourceGroup", "myRemediation",
            com.azure.core.util.Context.NONE);
    }
}
