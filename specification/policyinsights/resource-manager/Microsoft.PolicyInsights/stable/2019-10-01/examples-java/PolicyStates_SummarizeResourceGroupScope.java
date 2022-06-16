import com.azure.core.util.Context;

/** Samples for PolicyStates SummarizeForResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupScope.json
     */
    /**
     * Sample code: Summarize at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForResourceGroupWithResponse(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", null, null, null, null, Context.NONE);
    }
}
