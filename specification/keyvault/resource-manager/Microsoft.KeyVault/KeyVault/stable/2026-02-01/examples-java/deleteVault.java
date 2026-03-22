
/**
 * Samples for Vaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/deleteVault.json
     */
    /**
     * Sample code: Delete a vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void deleteAVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().deleteWithResponse("sample-resource-group", "sample-vault",
            com.azure.core.util.Context.NONE);
    }
}
