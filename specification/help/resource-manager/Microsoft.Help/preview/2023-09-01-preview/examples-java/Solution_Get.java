/** Samples for SolutionOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Get.json
     */
    /**
     * Sample code: Solution_Get.
     *
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionGet(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager
            .solutionOperations()
            .getWithResponse(
                "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
                "SolutionResource1",
                com.azure.core.util.Context.NONE);
    }
}
