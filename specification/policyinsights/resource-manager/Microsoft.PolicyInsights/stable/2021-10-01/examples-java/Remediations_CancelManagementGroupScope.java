import com.azure.core.util.Context;

/** Samples for Remediations CancelAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CancelManagementGroupScope.json
     */
    /**
     * Sample code: Cancel a remediation at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void cancelARemediationAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().cancelAtManagementGroupWithResponse("financeMg", "myRemediation", Context.NONE);
    }
}
