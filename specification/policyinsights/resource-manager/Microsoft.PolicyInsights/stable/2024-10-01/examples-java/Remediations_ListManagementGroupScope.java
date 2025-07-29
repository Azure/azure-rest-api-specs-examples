
/**
 * Samples for Remediations ListForManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * Remediations_ListManagementGroupScope.json
     */
    /**
     * Sample code: List remediations at management group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        listRemediationsAtManagementGroupScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().listForManagementGroup("financeMg", null, null, com.azure.core.util.Context.NONE);
    }
}
