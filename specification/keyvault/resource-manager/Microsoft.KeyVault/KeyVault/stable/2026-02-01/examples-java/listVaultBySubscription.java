
/**
 * Samples for Vaults ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listVaultBySubscription.json
     */
    /**
     * Sample code: List vaults in the specified subscription.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        listVaultsInTheSpecifiedSubscription(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().listBySubscription(1, com.azure.core.util.Context.NONE);
    }
}
