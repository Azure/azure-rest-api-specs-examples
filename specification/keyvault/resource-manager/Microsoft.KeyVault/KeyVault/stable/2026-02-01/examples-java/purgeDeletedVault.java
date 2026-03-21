
/**
 * Samples for Vaults PurgeDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/purgeDeletedVault.json
     */
    /**
     * Sample code: Purge a deleted vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void purgeADeletedVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().purgeDeleted("sample-vault", "westus", com.azure.core.util.Context.NONE);
    }
}
