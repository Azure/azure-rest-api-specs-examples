import com.azure.core.util.Context;

/** Samples for Remediations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListResourceGroupScope.json
     */
    /**
     * Sample code: List remediations at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listRemediationsAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().listByResourceGroup("myResourceGroup", null, null, Context.NONE);
    }
}
