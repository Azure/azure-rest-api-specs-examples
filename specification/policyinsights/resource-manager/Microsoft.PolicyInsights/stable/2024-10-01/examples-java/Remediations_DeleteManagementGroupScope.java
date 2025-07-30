
/**
 * Samples for Remediations DeleteAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * Remediations_DeleteManagementGroupScope.json
     */
    /**
     * Sample code: Delete remediation at management group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void deleteRemediationAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().deleteAtManagementGroupWithResponse("financeMg", "storageRemediation",
            com.azure.core.util.Context.NONE);
    }
}
