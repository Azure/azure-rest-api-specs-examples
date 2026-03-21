
/**
 * Samples for Vaults GetDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/getDeletedVault.json
     */
    /**
     * Sample code: Retrieve a deleted vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void retrieveADeletedVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().getDeletedWithResponse("sample-vault", "westus",
            com.azure.core.util.Context.NONE);
    }
}
