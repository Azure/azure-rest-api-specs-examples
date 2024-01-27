
/** Samples for Vaults GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/getVault.json
     */
    /**
     * Sample code: Retrieve a vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().getByResourceGroupWithResponse("sample-resource-group",
            "sample-vault", com.azure.core.util.Context.NONE);
    }
}
