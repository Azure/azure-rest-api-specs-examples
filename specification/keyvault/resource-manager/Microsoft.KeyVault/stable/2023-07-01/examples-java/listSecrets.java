
/**
 * Samples for Secrets List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/listSecrets.json
     */
    /**
     * Sample code: List secrets in the vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecretsInTheVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getSecrets().list("sample-group", "sample-vault", null,
            com.azure.core.util.Context.NONE);
    }
}
