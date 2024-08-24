
/**
 * Samples for Keys ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/listKeyVersions.json
     */
    /**
     * Sample code: List key versions in the vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKeyVersionsInTheVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getKeys().listVersions("sample-group", "sample-vault-name",
            "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
