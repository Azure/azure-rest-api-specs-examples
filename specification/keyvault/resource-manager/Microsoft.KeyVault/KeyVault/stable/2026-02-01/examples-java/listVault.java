
/**
 * Samples for Vaults List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listVault.json
     */
    /**
     * Sample code: List vaults in the specified subscription.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        listVaultsInTheSpecifiedSubscription(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().list(1, com.azure.core.util.Context.NONE);
    }
}
