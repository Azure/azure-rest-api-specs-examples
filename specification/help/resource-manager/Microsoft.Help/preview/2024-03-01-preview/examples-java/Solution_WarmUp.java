
/**
 * Samples for SolutionOperation WarmUp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Solution_WarmUp.json
     */
    /**
     * Sample code: Solution_WarmUp.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionWarmUp(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.solutionOperations().warmUpWithResponse(
            "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
            "SolutionResourceName1", null, com.azure.core.util.Context.NONE);
    }
}
