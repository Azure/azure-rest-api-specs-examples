
/**
 * Samples for Remediations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/
     * Remediations_GetResourceGroupScope.json
     */
    /**
     * Sample code: Get remediation at resource group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        getRemediationAtResourceGroupScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().getByResourceGroupWithResponse("myResourceGroup", "storageRemediation",
            com.azure.core.util.Context.NONE);
    }
}
