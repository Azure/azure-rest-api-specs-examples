import com.azure.core.util.Context;

/** Samples for Vaults GetDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/getDeletedVault.json
     */
    /**
     * Sample code: Retrieve a deleted vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveADeletedVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getVaults()
            .getDeletedWithResponse("sample-vault", "westus", Context.NONE);
    }
}
