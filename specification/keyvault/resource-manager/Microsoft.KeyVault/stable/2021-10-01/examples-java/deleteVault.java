import com.azure.core.util.Context;

/** Samples for Vaults Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/deleteVault.json
     */
    /**
     * Sample code: Delete a vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getVaults()
            .deleteWithResponse("sample-resource-group", "sample-vault", Context.NONE);
    }
}
