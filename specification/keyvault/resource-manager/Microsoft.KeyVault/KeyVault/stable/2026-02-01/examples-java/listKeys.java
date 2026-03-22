
/**
 * Samples for Keys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listKeys.json
     */
    /**
     * Sample code: List keys in the vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listKeysInTheVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getKeys().list("sample-group", "sample-vault-name", com.azure.core.util.Context.NONE);
    }
}
