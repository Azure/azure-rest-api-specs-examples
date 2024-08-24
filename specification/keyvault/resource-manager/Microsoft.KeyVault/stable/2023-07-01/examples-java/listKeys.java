
/**
 * Samples for Keys List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/listKeys.json
     */
    /**
     * Sample code: List keys in the vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKeysInTheVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getKeys().list("sample-group", "sample-vault-name",
            com.azure.core.util.Context.NONE);
    }
}
