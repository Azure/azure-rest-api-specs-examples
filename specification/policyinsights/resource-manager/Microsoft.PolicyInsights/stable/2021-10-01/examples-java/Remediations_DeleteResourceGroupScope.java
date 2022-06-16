import com.azure.core.util.Context;

/** Samples for Remediations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_DeleteResourceGroupScope.json
     */
    /**
     * Sample code: Delete remediation at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void deleteRemediationAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().deleteWithResponse("myResourceGroup", "storageRemediation", Context.NONE);
    }
}
