import com.azure.core.util.Context;

/** Samples for Remediations GetAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetManagementGroupScope.json
     */
    /**
     * Sample code: Get remediation at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getRemediationAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().getAtManagementGroupWithResponse("financeMg", "storageRemediation", Context.NONE);
    }
}
