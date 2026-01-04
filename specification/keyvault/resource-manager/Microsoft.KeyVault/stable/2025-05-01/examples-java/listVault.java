
/**
 * Samples for Vaults List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/listVault.json
     */
    /**
     * Sample code: List vaults in the specified subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVaultsInTheSpecifiedSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().list(1, com.azure.core.util.Context.NONE);
    }
}
