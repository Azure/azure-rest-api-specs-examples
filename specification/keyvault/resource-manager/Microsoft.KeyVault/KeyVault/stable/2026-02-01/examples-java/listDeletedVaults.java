
/**
 * Samples for Vaults ListDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listDeletedVaults.json
     */
    /**
     * Sample code: List deleted vaults in the specified subscription.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        listDeletedVaultsInTheSpecifiedSubscription(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().listDeleted(com.azure.core.util.Context.NONE);
    }
}
