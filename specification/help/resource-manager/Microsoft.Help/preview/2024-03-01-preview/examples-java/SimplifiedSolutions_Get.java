
/**
 * Samples for SimplifiedSolutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SimplifiedSolutions_Get.
     * json
     */
    /**
     * Sample code: Solution_Get.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionGet(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.simplifiedSolutions().getWithResponse(
            "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
            "simplifiedSolutionsResourceName1", com.azure.core.util.Context.NONE);
    }
}
