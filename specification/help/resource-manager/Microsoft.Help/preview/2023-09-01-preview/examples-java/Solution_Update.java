import com.azure.resourcemanager.selfhelp.models.SolutionResource;

/** Samples for SolutionOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Update.json
     */
    /**
     * Sample code: Solution_Update.
     *
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionUpdate(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        SolutionResource resource =
            manager
                .solutionOperations()
                .getWithResponse(
                    "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
                    "SolutionResourceName1",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
