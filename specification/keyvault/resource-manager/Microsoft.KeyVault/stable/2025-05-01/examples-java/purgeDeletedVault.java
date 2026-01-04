
/**
 * Samples for Vaults PurgeDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/purgeDeletedVault.json
     */
    /**
     * Sample code: Purge a deleted vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void purgeADeletedVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().purgeDeleted("sample-vault", "westus",
            com.azure.core.util.Context.NONE);
    }
}
