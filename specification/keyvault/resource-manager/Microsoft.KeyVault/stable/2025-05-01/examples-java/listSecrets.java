
/**
 * Samples for Secrets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/listSecrets.json
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
