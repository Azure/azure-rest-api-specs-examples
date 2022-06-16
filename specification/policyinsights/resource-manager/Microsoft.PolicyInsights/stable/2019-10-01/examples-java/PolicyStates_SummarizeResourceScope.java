import com.azure.core.util.Context;

/** Samples for PolicyStates SummarizeForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceScope.json
     */
    /**
     * Sample code: Summarize at resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForResourceWithResponse(
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault",
                2,
                null,
                null,
                null,
                Context.NONE);
    }
}
