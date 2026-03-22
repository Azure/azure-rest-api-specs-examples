
/**
 * Samples for Keys ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listKeyVersions.json
     */
    /**
     * Sample code: List key versions in the vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listKeyVersionsInTheVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getKeys().listVersions("sample-group", "sample-vault-name", "sample-key-name",
            com.azure.core.util.Context.NONE);
    }
}
