
/**
 * Samples for Remediations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * Remediations_DeleteResourceGroupScope.json
     */
    /**
     * Sample code: Delete remediation at resource group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        deleteRemediationAtResourceGroupScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().deleteByResourceGroupWithResponse("myResourceGroup", "storageRemediation",
            com.azure.core.util.Context.NONE);
    }
}
